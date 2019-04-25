package main

import (
	"io"
	"net/http"
	"time"
	"userService/pkg/pb"
	"userService/pkg/userservice"

	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/sd"
	consulsd "github.com/go-kit/kit/sd/consul"
	"github.com/go-kit/kit/sd/lb"
	"github.com/hashicorp/consul/api"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

func main() {
	logrus.SetLevel(logrus.DebugLevel)

	consulClient, err := api.NewClient(api.DefaultConfig())
	if err != nil {
		logrus.Fatal(err)
	}
	client := consulsd.NewClient(consulClient)
	log := &logger{}
	var (
		tags        = []string{}
		passingOnly = true
		instancer   = consulsd.NewInstancer(client, log, "userService", tags, passingOnly)
		endpoints   userservice.UserEndpoints
	)
	{
		factory := userserviceFactory(userservice.MakeLoginEndpoint)
		endpointer := sd.NewEndpointer(instancer, factory, log)
		balancer := lb.NewRoundRobin(endpointer)
		retry := lb.Retry(3, 5000*time.Millisecond, balancer)
		endpoints.LoginEndpoint = retry
	}
	{
		factory := userserviceFactory(userservice.MakeRegisterEndpoint)
		endpointer := sd.NewEndpointer(instancer, factory, log)
		balancer := lb.NewRoundRobin(endpointer)
		retry := lb.Retry(3, 500*time.Millisecond, balancer)
		endpoints.RegisterEndpoint = retry
	}
	{
		factory := userserviceFactory(userservice.MakeGetPermissionsEndpoint)
		endpointer := sd.NewEndpointer(instancer, factory, log)
		balancer := lb.NewRoundRobin(endpointer)
		retry := lb.Retry(3, 500*time.Millisecond, balancer)
		endpoints.GetPermissionsEndpoint = retry
	}
	{
		factory := userserviceFactory(userservice.MakeCheckPermissionEndpoint)
		endpointer := sd.NewEndpointer(instancer, factory, log)
		balancer := lb.NewRoundRobin(endpointer)
		retry := lb.Retry(3, 500*time.Millisecond, balancer)
		endpoints.CheckPermissionEndpoint = retry
	}
	{
		factory := userserviceFactory(userservice.MakeAddRoutesEndpoint)
		endpointer := sd.NewEndpointer(instancer, factory, log)
		balancer := lb.NewRoundRobin(endpointer)
		retry := lb.Retry(3, 500*time.Millisecond, balancer)
		endpoints.AddRoutesEndpoint = retry
	}
	{
		factory := userserviceFactory(userservice.MakeListRoutesEndpoint)
		endpointer := sd.NewEndpointer(instancer, factory, log)
		balancer := lb.NewRoundRobin(endpointer)
		retry := lb.Retry(3, 500*time.Millisecond, balancer)
		endpoints.ListRoutesEndpoint = retry
	}

	userHandler := userservice.NewHttpHandler(&endpoints)
	http.ListenAndServe(":8080", userHandler)
}

func userserviceFactory(makeEndpoint func(pb.UserServer) endpoint.Endpoint) sd.Factory {
	return func(instance string) (endpoint endpoint.Endpoint, closer io.Closer, e error) {
		conn, err := grpc.Dial(instance, grpc.WithInsecure())
		if err != nil {
			return nil, nil, err
		}
		service := userservice.NewGrpcClient(conn)
		return makeEndpoint(service), conn, nil
	}
}

type logger struct{}

func (l *logger) Log(args ...interface{}) error {
	logrus.Infoln(args...)
	return nil
}
