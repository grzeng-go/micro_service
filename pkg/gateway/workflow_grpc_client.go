package gateway

import (
	"context"
	"userService/pkg/pb"

	"github.com/go-kit/kit/endpoint"
	grpctransport "github.com/go-kit/kit/transport/grpc"
	kitgrpc "github.com/go-kit/kit/transport/grpc"
	"google.golang.org/grpc"
)

type WorkflowEndpoints struct {
	StartEndpoint      endpoint.Endpoint
	ListTaskEndpoint   endpoint.Endpoint
	HandleTaskEndpoint endpoint.Endpoint
}

func (w *WorkflowEndpoints) Start(ctx context.Context, in *pb.StartWorkflowRequest) (*pb.StartWorkflowReply, error) {
	res, err := w.StartEndpoint(ctx, in)
	if err != nil {
		return nil, err
	}
	return res.(*pb.StartWorkflowReply), nil
}

func (w *WorkflowEndpoints) ListTask(ctx context.Context, in *pb.ListTaskRequest) (*pb.ListTaskReply, error) {
	res, err := w.StartEndpoint(ctx, in)
	if err != nil {
		return nil, err
	}
	return res.(*pb.ListTaskReply), nil
}

func (w *WorkflowEndpoints) HandleTask(ctx context.Context, in *pb.HandleTaskRequest) (*pb.HandleTaskReply, error) {
	res, err := w.StartEndpoint(ctx, in)
	if err != nil {
		return nil, err
	}
	return res.(*pb.HandleTaskReply), nil
}

func NewWorkflowGRPCClient(conn *grpc.ClientConn, tracer kitgrpc.ClientOption) *WorkflowEndpoints {
	endpoints := new(WorkflowEndpoints)
	options := make([]kitgrpc.ClientOption, 0)
	if tracer != nil {
		options = append(options, tracer)
	}

	{
		e := grpctransport.NewClient(
			conn,
			"pb.Workflow",
			"ListTask",
			encodeRequest,
			decodeResponse,
			pb.ListTaskReply{},
			options...,
		).Endpoint()
		endpoints.ListTaskEndpoint = e
	}

	{
		e := grpctransport.NewClient(
			conn,
			"pb.Workflow",
			"HandleTask",
			encodeRequest,
			decodeResponse,
			pb.HandleTaskReply{},
			options...,
		).Endpoint()
		endpoints.HandleTaskEndpoint = e
	}

	{
		e := grpctransport.NewClient(
			conn,
			"pb.Workflow",
			"Start",
			encodeRequest,
			decodeResponse,
			pb.StartWorkflowReply{},
			options...,
		).Endpoint()
		endpoints.StartEndpoint = e
	}

	return endpoints
}
