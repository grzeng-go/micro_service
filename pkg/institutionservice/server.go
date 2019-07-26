package institutionservice

import (
	"context"
	"userService/pkg/pb"

	grpctransport "github.com/go-kit/kit/transport/grpc"
)

type server struct {
	tnxHisDownload                       grpctransport.Handler
	getTfrTrnLogs                        grpctransport.Handler
	getTfrTrnLog                         grpctransport.Handler
	downloadTfrTrnLogs                   grpctransport.Handler
	listGroupsHandler                    grpctransport.Handler
	listInstitutionsHandler              grpctransport.Handler
	addInstitutionHandler                grpctransport.Handler
	GetInstitutionByIdHandler            grpctransport.Handler
	SaveInstitutionFeeControlCashHandler grpctransport.Handler
}

func (s *server) SaveInstitutionFeeControlCash(ctx context.Context, in *pb.SaveInstitutionFeeControlCashRequest) (*pb.SaveInstitutionFeeControlCashReply, error) {

	_, res, err := s.SaveInstitutionFeeControlCashHandler.ServeGRPC(ctx, in)
	if err != nil {
		return nil, err
	}
	return res.(*pb.SaveInstitutionFeeControlCashReply), nil
}

func (s *server) GetInstitutionById(ctx context.Context, in *pb.GetInstitutionByIdRequest) (*pb.GetInstitutionByIdReply, error) {
	_, res, err := s.GetInstitutionByIdHandler.ServeGRPC(ctx, in)
	if err != nil {
		return nil, err
	}
	return res.(*pb.GetInstitutionByIdReply), nil
}

func (s *server) TnxHisDownload(ctx context.Context, in *pb.InstitutionTnxHisDownloadReq) (*pb.InstitutionTnxHisDownloadResp, error) {
	_, res, err := s.tnxHisDownload.ServeGRPC(ctx, in)
	if err != nil {
		return nil, err
	}
	return res.(*pb.InstitutionTnxHisDownloadResp), nil
}

func (s *server) GetTfrTrnLogs(ctx context.Context, in *pb.GetTfrTrnLogsReq) (*pb.GetTfrTrnLogsResp, error) {
	_, res, err := s.getTfrTrnLogs.ServeGRPC(ctx, in)
	if err != nil {
		return nil, err
	}
	return res.(*pb.GetTfrTrnLogsResp), nil
}

func (s *server) GetTfrTrnLog(ctx context.Context, in *pb.GetTfrTrnLogReq) (*pb.GetTfrTrnLogResp, error) {
	_, res, err := s.getTfrTrnLog.ServeGRPC(ctx, in)
	if err != nil {
		return nil, err
	}
	return res.(*pb.GetTfrTrnLogResp), nil
}

func (s *server) DownloadTfrTrnLogs(ctx context.Context, in *pb.DownloadTfrTrnLogsReq) (*pb.DownloadTfrTrnLogsResp, error) {
	_, res, err := s.downloadTfrTrnLogs.ServeGRPC(ctx, in)
	if err != nil {
		return nil, err
	}
	return res.(*pb.DownloadTfrTrnLogsResp), nil
}

func (s *server) ListGroups(ctx context.Context, in *pb.ListGroupsRequest) (*pb.ListInstitutionsReply, error) {
	_, res, err := s.listGroupsHandler.ServeGRPC(ctx, in)
	if err != nil {
		return nil, err
	}
	reply, ok := res.(*pb.ListInstitutionsReply)
	if !ok {
		return nil, ErrReplyTypeInvalid
	}
	return reply, nil
}

func (s *server) ListInstitutions(ctx context.Context, in *pb.ListInstitutionsRequest) (*pb.ListInstitutionsReply, error) {
	_, res, err := s.listInstitutionsHandler.ServeGRPC(ctx, in)
	if err != nil {
		return nil, err
	}
	reply, ok := res.(*pb.ListInstitutionsReply)
	if !ok {
		return nil, ErrReplyTypeInvalid
	}
	return reply, nil
}

func (s *server) SaveInstitution(ctx context.Context, in *pb.SaveInstitutionRequest) (*pb.SaveInstitutionReply, error) {
	_, res, err := s.addInstitutionHandler.ServeGRPC(ctx, in)
	if err != nil {
		return nil, err
	}
	reply, ok := res.(*pb.SaveInstitutionReply)
	if !ok {
		return nil, ErrReplyTypeInvalid
	}
	return reply, nil

}

func New(tracer grpctransport.ServerOption) pb.InstitutionServer {
	svc := new(service)
	svr := new(server)
	options := make([]grpctransport.ServerOption, 0)
	if tracer != nil {
		options = append(options, tracer)
	}

	{
		e := MakeGetInstitutionByIdEndpoint(svc)
		svr.GetInstitutionByIdHandler = grpctransport.NewServer(
			e,
			grpcDecode,
			grpcEncode,
			options...,
		)
	}

	{
		e := MakeTnxHisDownloadEndpoint(svc)
		svr.tnxHisDownload = grpctransport.NewServer(
			e,
			grpcDecode,
			grpcEncode,
			options...,
		)
	}
	{
		e := MakeGetTfrTrnLogsEndpoint(svc)
		svr.getTfrTrnLogs = grpctransport.NewServer(
			e,
			grpcDecode,
			grpcEncode,
			options...,
		)
	}
	{
		e := MakeGetTfrTrnLogEndpoint(svc)
		svr.getTfrTrnLog = grpctransport.NewServer(
			e,
			grpcDecode,
			grpcEncode,
			options...,
		)
	}
	{
		e := MakeDownloadTfrTrnLogsEndpoint(svc)
		svr.downloadTfrTrnLogs = grpctransport.NewServer(
			e,
			grpcDecode,
			grpcEncode,
			options...,
		)
	}

	{
		e := MakeListGroupsEndpoint(svc)
		svr.listGroupsHandler = grpctransport.NewServer(
			e,
			grpcDecode,
			grpcEncode,
			options...,
		)
	}

	{
		e := MakeListInstitutionsEndpoint(svc)
		svr.listInstitutionsHandler = grpctransport.NewServer(
			e,
			grpcDecode,
			grpcEncode,
			options...,
		)
	}

	{
		e := MakeSaveInstitutionEndpoint(svc)
		svr.addInstitutionHandler = grpctransport.NewServer(
			e,
			grpcDecode,
			grpcEncode,
			options...,
		)
	}

	{
		e := MakeSaveInstitutionFeeControlCashEndpoint(svc)
		svr.SaveInstitutionFeeControlCashHandler = grpctransport.NewServer(
			e,
			grpcDecode,
			grpcEncode,
			options...,
		)
	}

	return svr
}
