package termservice

import (
	"context"
	"userService/pkg/kit"
	"userService/pkg/pb"

	grpctransport "github.com/go-kit/kit/transport/grpc"
)

type server struct {
	ListTermInfoHandler            grpctransport.Handler
	SaveTermHandler                grpctransport.Handler
	ListTermRiskHandler            grpctransport.Handler
	ListTermActivationStateHandler grpctransport.Handler
}

func (s *server) ListTermActivationState(ctx context.Context, in *pb.ListTermActivationStateRequest) (*pb.ListTermActivationStateReply, error) {
	_, res, err := s.ListTermActivationStateHandler.ServeGRPC(ctx, in)
	if err != nil {
		return nil, err
	}
	return res.(*pb.ListTermActivationStateReply), nil
}

func (s *server) ListTermInfo(ctx context.Context, in *pb.ListTermInfoRequest) (*pb.ListTermInfoReply, error) {
	_, res, err := s.ListTermInfoHandler.ServeGRPC(ctx, in)
	if err != nil {
		return nil, err
	}
	return res.(*pb.ListTermInfoReply), nil
}

func (s *server) SaveTerm(ctx context.Context, in *pb.SaveTermRequest) (*pb.SaveTermReply, error) {
	_, res, err := s.SaveTermHandler.ServeGRPC(ctx, in)
	if err != nil {
		return nil, err
	}
	return res.(*pb.SaveTermReply), nil
}

func (s *server) ListTermRisk(ctx context.Context, in *pb.ListTermRiskRequest) (*pb.ListTermRiskReply, error) {
	_, res, err := s.ListTermRiskHandler.ServeGRPC(ctx, in)
	if err != nil {
		return nil, err
	}
	return res.(*pb.ListTermRiskReply), nil
}

func New(tracer grpctransport.ServerOption) pb.TermServer {
	svr := &server{}
	svc := &service{}
	options := make([]grpctransport.ServerOption, 0)
	if tracer != nil {
		options = append(options, tracer)
	}

	{
		endpoint := MakeListTermInfoEndpoint(svc)
		endpoint = kit.LogginMiddleware(endpoint)
		svr.ListTermInfoHandler = grpctransport.NewServer(
			endpoint,
			kit.DecodeRequest,
			kit.EncodeResponse,
			options...,
		)
	}

	{
		endpoint := MakeSaveTermEndpoint(svc)
		endpoint = kit.LogginMiddleware(endpoint)
		svr.SaveTermHandler = grpctransport.NewServer(
			endpoint,
			kit.DecodeRequest,
			kit.EncodeResponse,
			options...,
		)
	}

	{
		endpoint := MakeListTermRiskEndpoint(svc)
		endpoint = kit.LogginMiddleware(endpoint)
		svr.ListTermRiskHandler = grpctransport.NewServer(
			endpoint,
			kit.DecodeRequest,
			kit.EncodeResponse,
			options...,
		)
	}

	{
		endpoint := MakeListTermActivationStateEndpoint(svc)
		endpoint = kit.LogginMiddleware(endpoint)
		svr.ListTermActivationStateHandler = grpctransport.NewServer(
			endpoint,
			kit.DecodeRequest,
			kit.EncodeResponse,
			options...,
		)
	}

	return svr
}
