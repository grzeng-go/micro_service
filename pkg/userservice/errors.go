package userservice

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	ErrUserExists              = status.Error(codes.AlreadyExists, "user is exists")
	ErrUserNamePasswordEmpty   = status.Error(codes.NotFound, "empty username or password")
	ErrWrongUserNameOrPassword = status.Error(codes.InvalidArgument, "wrong username or password")
	ErrUserNotFound            = status.Error(codes.NotFound, "user not found")
	ErrReplyTypeInvalid        = status.Error(codes.Internal, "reply type invalid")
	ErrRequestTypeInvalid      = status.Error(codes.Internal, "request type invalid")
	ErrUnexpectedSigningMethod = status.Error(codes.PermissionDenied, "unexpected signing method")
	ErrInvalidParams           = status.Error(codes.InvalidArgument, "invalid params")
	ErrPolicyExists            = status.Error(codes.AlreadyExists, "policy exists")
	ErrPolicyNotFound          = status.Error(codes.NotFound, "policy not found")
	ErrRoleExists              = status.Error(codes.AlreadyExists, "role exists")
	ErrRoleNotFound            = status.Error(codes.NotFound, "role not found")
	ErrRouteExists             = status.Error(codes.AlreadyExists, "routes exists")
	ErrPermissionExists        = status.Error(codes.AlreadyExists, "permission exists")
	ErrPermissionNotFound      = status.Error(codes.NotFound, "permission not found")
	ErrRouteNotFound           = status.Error(codes.NotFound, "route not found")
)

const (
	InvalidParam  = "InvalidParamError"
	AlreadyExists = "AlreadyExistsError"
	NotFound      = "NotFoundError"
)
