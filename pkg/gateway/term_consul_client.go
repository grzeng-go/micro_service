package gateway

import (
	"io"
	"os"
	"time"
	"userService/pkg/pb"
	"userService/pkg/termservice"

	"github.com/afex/hystrix-go/hystrix"
	"github.com/go-kit/kit/circuitbreaker"
	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/sd"
	"github.com/go-kit/kit/sd/lb"
	"github.com/go-kit/kit/tracing/zipkin"
	stdzipkin "github.com/openzipkin/zipkin-go"
	zipkinhttp "github.com/openzipkin/zipkin-go/reporter/http"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

func GetTermEndpoints(instancer sd.Instancer, log log.Logger) *TermEndpoints {
	var endpoints TermEndpoints

	hystrix.ConfigureCommand(userbreaker, hystrix.CommandConfig{
		MaxConcurrentRequests: 1000,
		Timeout:               10000,
		ErrorPercentThreshold: 25,
		SleepWindow:           10000,
	})
	userBreaker := circuitbreaker.Hystrix(userbreaker)

	{
		factory := termServiceFactory(termservice.MakeListTermInfoEndpoint)
		endpointer := sd.NewEndpointer(instancer, factory, log)
		balancer := lb.NewRoundRobin(endpointer)
		retry := lb.Retry(3, 5000*time.Millisecond, balancer)
		retry = userBreaker(retry)
		endpoints.ListTermInfoEndpoint = retry
	}

	return &endpoints
}

func termServiceFactory(makeEndpoint func(server pb.TermServer) endpoint.Endpoint) sd.Factory {
	return func(instance string) (endpoint endpoint.Endpoint, closer io.Closer, e error) {
		conn, err := grpc.Dial(instance, grpc.WithInsecure())
		if err != nil {
			return nil, nil, err
		}
		addr := os.Getenv("ZIPKIN_ADDR")
		if addr == "" {
			addr = "127.0.0.1:9411"
		}

		localEndpoint, _ := stdzipkin.NewEndpoint("merchant", "localhost:9411")
		reporter := zipkinhttp.NewReporter("http://" + addr + "/api/v2/spans")
		stdTracer, err := stdzipkin.NewTracer(
			reporter,
			stdzipkin.WithLocalEndpoint(localEndpoint),
		)
		if err != nil {
			logrus.Errorln(err)
		}
		var service *TermEndpoints
		if stdTracer == nil {
			service = NewTermServiceClient(conn, nil)
		} else {
			tracer := zipkin.GRPCClientTrace(stdTracer)
			service = NewTermServiceClient(conn, tracer)
		}

		return makeEndpoint(service), conn, nil
	}
}