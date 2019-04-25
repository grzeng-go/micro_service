package userservice

import (
	"context"
	"userService/pkg/pb"

	"github.com/go-kit/kit/endpoint"
	"google.golang.org/grpc"

	grpctransport "github.com/go-kit/kit/transport/grpc"
)

type UserEndpoints struct {
	LoginEndpoint                         endpoint.Endpoint
	GetPermissionsEndpoint                endpoint.Endpoint
	CheckPermissionEndpoint               endpoint.Endpoint
	RegisterEndpoint                      endpoint.Endpoint
	AddPermissionForRoleEndpoint          endpoint.Endpoint
	AddRoleForRoleEndpoint                endpoint.Endpoint
	CreateRoleEndpoint                    endpoint.Endpoint
	AddRoleForUserEndpoint                endpoint.Endpoint
	AddRoutesEndpoint                     endpoint.Endpoint
	ListRouteEndpoint                     endpoint.Endpoint
	CreatePermissionEndpoint              endpoint.Endpoint
	UpdatePermissionEndpoint              endpoint.Endpoint
	AddRouteForPermissionEndpoint         endpoint.Endpoint
	RemoveRouteForPermissionEndpoint      endpoint.Endpoint
	RemovePermissionEndpoint              endpoint.Endpoint
	ListPermissionsEndpoint               endpoint.Endpoint
	AddPermissionForPermissionEndpoint    endpoint.Endpoint
	RemovePermissionForPermissionEndpoint endpoint.Endpoint
	ListRoleEndpoint                      endpoint.Endpoint
	UpdateRoleEndpoint                    endpoint.Endpoint
	RemovePermissionForRoleEndpoint       endpoint.Endpoint
	RemoveRoleForRoleEndpoint             endpoint.Endpoint
	RemoveRoleEndpoint                    endpoint.Endpoint
	ListUsersEndpoint                     endpoint.Endpoint
	UpdateUserEndpoint                    endpoint.Endpoint
	AddPermissionForUserEndpoint          endpoint.Endpoint
	RemovePermissionForUserEndpoint       endpoint.Endpoint
	RemoveRoleForUserEndpoint             endpoint.Endpoint
}

func NewGrpcClient(conn *grpc.ClientConn) *UserEndpoints {
	endpoints := new(UserEndpoints)
	{
		endpoint := grpctransport.NewClient(
			conn,
			"pb.User",
			"Login",
			encodeRequest,
			decodeResponse,
			pb.LoginReply{},
		).Endpoint()
		endpoints.LoginEndpoint = endpoint
	}

	{
		endpoint := grpctransport.NewClient(
			conn,
			"pb.User",
			"GetPermissions",
			encodeRequest,
			decodeResponse,
			pb.GetPermissionsReply{},
		).Endpoint()
		endpoints.GetPermissionsEndpoint = endpoint
	}

	{
		endpoint := grpctransport.NewClient(
			conn,
			"pb.User",
			"CheckPermission",
			encodeRequest,
			decodeResponse,
			pb.CheckPermissionReply{},
		).Endpoint()
		endpoints.CheckPermissionEndpoint = endpoint
	}

	{
		endpoint := grpctransport.NewClient(
			conn,
			"pb.User",
			"Register",
			encodeRequest,
			decodeResponse,
			pb.RegisterReply{},
		).Endpoint()
		endpoints.RegisterEndpoint = endpoint
	}
	return endpoints
}
func (u *UserEndpoints) Login(ctx context.Context, in *pb.LoginRequest) (*pb.LoginReply, error) {
	res, err := u.LoginEndpoint(ctx, in)
	if err != nil {
		return nil, err
	}
	reply, ok := res.(*pb.LoginReply)
	if !ok {
		return nil, ErrReplyTypeInvalid
	}
	return reply, nil
}

func (u *UserEndpoints) GetPermissions(ctx context.Context, in *pb.GetPermissionsRequest) (*pb.GetPermissionsReply, error) {
	res, err := u.GetPermissionsEndpoint(ctx, in)
	if err != nil {
		return nil, err
	}
	reply, ok := res.(*pb.GetPermissionsReply)
	if !ok {
		return nil, ErrReplyTypeInvalid
	}
	return reply, nil
}

func (u *UserEndpoints) CheckPermission(ctx context.Context, in *pb.CheckPermissionRequest) (*pb.CheckPermissionReply, error) {
	res, err := u.CheckPermissionEndpoint(ctx, in)
	if err != nil {
		return nil, err
	}
	reply, ok := res.(*pb.CheckPermissionReply)
	if !ok {
		return nil, ErrReplyTypeInvalid
	}
	return reply, nil
}

func (u *UserEndpoints) Register(ctx context.Context, in *pb.RegisterRequest) (*pb.RegisterReply, error) {
	res, err := u.RegisterEndpoint(ctx, in)
	if err != nil {
		return nil, err
	}
	reply, ok := res.(*pb.RegisterReply)
	if !ok {
		return nil, ErrReplyTypeInvalid
	}
	return reply, nil
}

func (u *UserEndpoints) AddPermissionForRole(ctx context.Context, in *pb.AddPermissionForRoleRequest) (*pb.AddPermissionForRoleReply, error) {
	res, err := u.AddPermissionForRoleEndpoint(ctx, in)
	if err != nil {
		return nil, err
	}
	reply, ok := res.(*pb.AddPermissionForRoleReply)
	if !ok {
		return nil, ErrReplyTypeInvalid
	}
	return reply, nil
}

func (u *UserEndpoints) CreateRole(ctx context.Context, in *pb.CreateRoleRequest) (*pb.CreateRoleReply, error) {
	res, err := u.CreateRoleEndpoint(ctx, in)
	if err != nil {
		return nil, err
	}
	reply, ok := res.(*pb.CreateRoleReply)
	if !ok {
		return nil, ErrReplyTypeInvalid
	}
	return reply, nil
}

func (u *UserEndpoints) AddRoleForUser(ctx context.Context, in *pb.AddRoleForUserRequest) (*pb.AddRoleForUserReply, error) {
	res, err := u.AddRoleForUserEndpoint(ctx, in)
	if err != nil {
		return nil, err
	}
	reply, ok := res.(*pb.AddRoleForUserReply)
	if !ok {
		return nil, ErrReplyTypeInvalid
	}
	return reply, nil
}

func (u *UserEndpoints) AddRoutes(ctx context.Context, in *pb.AddRoutesRequest) (*pb.AddRoutesReply, error) {
	res, err := u.AddRoutesEndpoint(ctx, in)
	if err != nil {
		return nil, err
	}
	reply, ok := res.(*pb.AddRoutesReply)
	if !ok {
		return nil, ErrReplyTypeInvalid
	}
	return reply, nil
}

func (u *UserEndpoints) ListRoutes(ctx context.Context, in *pb.ListRoutesRequest) (*pb.ListRoutesReply, error) {
	res, err := u.ListRouteEndpoint(ctx, in)
	if err != nil {
		return nil, err
	}
	reply, ok := res.(*pb.ListRoutesReply)
	if !ok {
		return nil, ErrReplyTypeInvalid
	}
	return reply, nil
}

func (u *UserEndpoints) CreatePermission(ctx context.Context, in *pb.CreatePermissionRequest) (*pb.CreatePermissionReply, error) {
	res, err := u.CreatePermissionEndpoint(ctx, in)
	if err != nil {
		return nil, err
	}
	reply, ok := res.(*pb.CreatePermissionReply)
	if !ok {
		return nil, ErrReplyTypeInvalid
	}
	return reply, nil
}

func (u *UserEndpoints) UpdatePermission(ctx context.Context, in *pb.UpdatePermissionRequest) (*pb.UpdatePermissionReply, error) {
	res, err := u.UpdatePermissionEndpoint(ctx, in)
	if err != nil {
		return nil, err
	}
	reply, ok := res.(*pb.UpdatePermissionReply)
	if !ok {
		return nil, ErrReplyTypeInvalid
	}
	return reply, nil
}

func (u *UserEndpoints) AddRouteForPermission(ctx context.Context, in *pb.AddRouteForPermissionRequest) (*pb.AddRouteForPermissionReply, error) {
	res, err := u.AddRouteForPermissionEndpoint(ctx, in)
	if err != nil {
		return nil, err
	}
	reply, ok := res.(*pb.AddRouteForPermissionReply)
	if !ok {
		return nil, ErrReplyTypeInvalid
	}
	return reply, nil
}

func (u *UserEndpoints) RemoveRouteForPermission(ctx context.Context, in *pb.RemoveRouteForPermissionRequest) (*pb.RemoveRouteForPermissionReply, error) {
	res, err := u.RemoveRouteForPermissionEndpoint(ctx, in)
	if err != nil {
		return nil, err
	}
	reply, ok := res.(*pb.RemoveRouteForPermissionReply)
	if !ok {
		return nil, ErrReplyTypeInvalid
	}
	return reply, nil
}

func (u *UserEndpoints) RemovePermission(ctx context.Context, in *pb.RemovePermissionRequest) (*pb.RemovePermissionReply, error) {
	res, err := u.RemovePermissionEndpoint(ctx, in)
	if err != nil {
		return nil, err
	}
	reply, ok := res.(*pb.RemovePermissionReply)
	if !ok {
		return nil, ErrReplyTypeInvalid
	}
	return reply, nil
}

func (u *UserEndpoints) ListPermissions(ctx context.Context, in *pb.ListPermissionsRequest) (*pb.ListPermissionsReply, error) {
	res, err := u.ListPermissionsEndpoint(ctx, in)
	if err != nil {
		return nil, err
	}
	reply, ok := res.(*pb.ListPermissionsReply)
	if !ok {
		return nil, ErrReplyTypeInvalid
	}
	return reply, nil
}

func (u *UserEndpoints) AddPermissionForPermission(ctx context.Context, in *pb.AddPermissionForPermissionRequest) (*pb.AddPermissionForPermissionReply, error) {
	res, err := u.AddPermissionForPermissionEndpoint(ctx, in)
	if err != nil {
		return nil, err
	}
	reply, ok := res.(*pb.AddPermissionForPermissionReply)
	if !ok {
		return nil, ErrReplyTypeInvalid
	}
	return reply, nil
}

func (u *UserEndpoints) RemovePermissionForPermission(ctx context.Context, in *pb.RemovePermissionForPermissionRequest) (*pb.RemovePermissionForPermissionReply, error) {
	res, err := u.RemovePermissionForPermissionEndpoint(ctx, in)
	if err != nil {
		return nil, err
	}
	reply, ok := res.(*pb.RemovePermissionForPermissionReply)
	if !ok {
		return nil, ErrReplyTypeInvalid
	}
	return reply, nil
}

func (u *UserEndpoints) ListRole(ctx context.Context, in *pb.ListRoleRequest) (*pb.ListRoleReply, error) {
	res, err := u.ListRoleEndpoint(ctx, in)
	if err != nil {
		return nil, err
	}
	reply, ok := res.(*pb.ListRoleReply)
	if !ok {
		return nil, ErrReplyTypeInvalid
	}
	return reply, nil
}

func (u *UserEndpoints) UpdateRole(ctx context.Context, in *pb.UpdateRoleRequest) (*pb.UpdateRoleReply, error) {
	res, err := u.UpdateRoleEndpoint(ctx, in)
	if err != nil {
		return nil, err
	}
	reply, ok := res.(*pb.UpdateRoleReply)
	if !ok {
		return nil, ErrReplyTypeInvalid
	}
	return reply, nil
}

func (u *UserEndpoints) RemovePermissionForRole(ctx context.Context, in *pb.RemovePermissionForRoleRequest) (*pb.RemovePermissionForRoleReply, error) {
	res, err := u.RemovePermissionForRoleEndpoint(ctx, in)
	if err != nil {
		return nil, err
	}
	reply, ok := res.(*pb.RemovePermissionForRoleReply)
	if !ok {
		return nil, ErrReplyTypeInvalid
	}
	return reply, nil
}

func (u *UserEndpoints) AddRoleForRole(ctx context.Context, in *pb.AddRoleForRoleRequest) (*pb.AddRoleForRoleReply, error) {
	res, err := u.AddRoleForRoleEndpoint(ctx, in)
	if err != nil {
		return nil, err
	}
	reply, ok := res.(*pb.AddRoleForRoleReply)
	if !ok {
		return nil, ErrReplyTypeInvalid
	}
	return reply, nil
}

func (u *UserEndpoints) RemoveRoleForRole(ctx context.Context, in *pb.RemoveRoleForRoleRequest) (*pb.RemoveRoleForRoleReply, error) {
	res, err := u.RemoveRoleForRoleEndpoint(ctx, in)
	if err != nil {
		return nil, err
	}
	reply, ok := res.(*pb.RemoveRoleForRoleReply)
	if !ok {
		return nil, ErrReplyTypeInvalid
	}
	return reply, nil
}

func (u *UserEndpoints) RemoveRole(ctx context.Context, in *pb.RemoveRoleRequest) (*pb.RemoveRoleReply, error) {
	res, err := u.RemoveRoleEndpoint(ctx, in)
	if err != nil {
		return nil, err
	}
	reply, ok := res.(*pb.RemoveRoleReply)
	if !ok {
		return nil, ErrReplyTypeInvalid
	}
	return reply, nil
}

func (u *UserEndpoints) ListUsers(ctx context.Context, in *pb.ListUsersRequest) (*pb.ListUsersReply, error) {
	res, err := u.ListUsersEndpoint(ctx, in)
	if err != nil {
		return nil, err
	}
	reply, ok := res.(*pb.ListUsersReply)
	if !ok {
		return nil, ErrReplyTypeInvalid
	}
	return reply, nil
}

func (u *UserEndpoints) UpdateUser(ctx context.Context, in *pb.UpdateUserRequest) (*pb.UpdateUserReply, error) {
	res, err := u.UpdateUserEndpoint(ctx, in)
	if err != nil {
		return nil, err
	}
	reply, ok := res.(*pb.UpdateUserReply)
	if !ok {
		return nil, ErrReplyTypeInvalid
	}
	return reply, nil
}

func (u *UserEndpoints) AddPermissionForUser(ctx context.Context, in *pb.AddPermissionForUserRequest) (*pb.AddPermissionForUserReply, error) {
	res, err := u.AddPermissionForUserEndpoint(ctx, in)
	if err != nil {
		return nil, err
	}
	reply, ok := res.(*pb.AddPermissionForUserReply)
	if !ok {
		return nil, ErrReplyTypeInvalid
	}
	return reply, nil
}

func (u *UserEndpoints) RemovePermissionForUser(ctx context.Context, in *pb.RemovePermissionForUserRequest) (*pb.RemovePermissionForUserReply, error) {
	res, err := u.RemovePermissionForUserEndpoint(ctx, in)
	if err != nil {
		return nil, err
	}
	reply, ok := res.(*pb.RemovePermissionForUserReply)
	if !ok {
		return nil, ErrReplyTypeInvalid
	}
	return reply, nil
}

func (u *UserEndpoints) RemoveRoleForUser(ctx context.Context, in *pb.RemoveRoleForUserRequest) (*pb.RemoveRoleForUserReply, error) {
	res, err := u.RemoveRoleForUserEndpoint(ctx, in)
	if err != nil {
		return nil, err
	}
	reply, ok := res.(*pb.RemoveRoleForUserReply)
	if !ok {
		return nil, ErrReplyTypeInvalid
	}
	return reply, nil
}
