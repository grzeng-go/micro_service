package gateway

import (
	"context"
	"encoding/json"
	"net/http"
	"userService/pkg/pb"
	"userService/pkg/userservice"

	"github.com/dgrijalva/jwt-go"
	stdjwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/go-kit/kit/sd/lb"
	httptransport "github.com/go-kit/kit/transport/http"
	"google.golang.org/grpc/metadata"
)

func RegisterUserHandler(engine *gin.Engine, endpoints *UserEndpoints) {
	engine.POST("/user/login", convertHttpHandlerToGinHandler(httptransport.NewServer(
		endpoints.LoginEndpoint,
		decodeHttpLoginRequest,
		encodeHttpResponse,
		httptransport.ServerErrorEncoder(errorEncoder),
	)))

	engine.POST("/user/register", convertHttpHandlerToGinHandler(httptransport.NewServer(
		endpoints.RegisterEndpoint,
		decodeHttpRegisterRequest,
		encodeHttpResponse,
		httptransport.ServerErrorEncoder(errorEncoder),
	)))

	engine.POST("/user/getPermissions",
		userservice.JwtMiddleware(keyFunc, stdjwt.SigningMethodHS256, userservice.UserClaimFactory),
		convertHttpHandlerToGinHandler(httptransport.NewServer(
			endpoints.GetPermissionsEndpoint,
			decodeHttpGetPermissionsRequest,
			encodeHttpResponse,
			httptransport.ServerErrorEncoder(errorEncoder),
			httptransport.ServerBefore(setUserInfoContext),
		)))

	engine.POST("/user/checkPermission",
		userservice.JwtMiddleware(keyFunc, stdjwt.SigningMethodHS256, userservice.UserClaimFactory),
		convertHttpHandlerToGinHandler(httptransport.NewServer(
			endpoints.CheckPermissionEndpoint,
			decodeHttpCheckPermissionRequest,
			encodeHttpResponse,
			httptransport.ServerErrorEncoder(errorEncoder),
			httptransport.ServerBefore(setUserInfoContext),
		)))

	engine.POST("/user/addRoutes",
		userservice.JwtMiddleware(keyFunc, stdjwt.SigningMethodHS256, userservice.UserClaimFactory),
		convertHttpHandlerToGinHandler(httptransport.NewServer(
			endpoints.AddRoutesEndpoint,
			decodeHttpAddRoutesRequest,
			encodeHttpResponse,
			httptransport.ServerErrorEncoder(errorEncoder),
		)))

	engine.POST("/user/listRoutes",
		userservice.JwtMiddleware(keyFunc, stdjwt.SigningMethodHS256, userservice.UserClaimFactory),
		convertHttpHandlerToGinHandler(httptransport.NewServer(
			endpoints.ListRoutesEndpoint,
			decodeHttpListRoutesRequest,
			encodeHttpResponse,
			httptransport.ServerErrorEncoder(errorEncoder),
		)))

	engine.POST("/user/createPermission",
		userservice.JwtMiddleware(keyFunc, stdjwt.SigningMethodHS256, userservice.UserClaimFactory),
		convertHttpHandlerToGinHandler(httptransport.NewServer(
			endpoints.CreatePermissionEndpoint,
			decodeHttpCreatePermissionRequest,
			encodeHttpResponse,
			httptransport.ServerErrorEncoder(errorEncoder),
		)))

	engine.POST("/user/updatePermission",
		userservice.JwtMiddleware(keyFunc, stdjwt.SigningMethodHS256, userservice.UserClaimFactory),
		convertHttpHandlerToGinHandler(httptransport.NewServer(
			endpoints.UpdatePermissionEndpoint,
			decodeHttpUpdatePermissionRequest,
			encodeHttpResponse,
			httptransport.ServerErrorEncoder(errorEncoder),
		)))

	engine.POST("/user/addRouteForPermission",
		userservice.JwtMiddleware(keyFunc, stdjwt.SigningMethodHS256, userservice.UserClaimFactory),
		convertHttpHandlerToGinHandler(httptransport.NewServer(
			endpoints.AddRouteForPermissionEndpoint,
			decodeHttpAddRouteForPermissionRequest,
			encodeHttpResponse,
			httptransport.ServerErrorEncoder(errorEncoder),
		)))

	engine.POST("/user/removeRouteForPermission",
		userservice.JwtMiddleware(keyFunc, stdjwt.SigningMethodHS256, userservice.UserClaimFactory),
		convertHttpHandlerToGinHandler(httptransport.NewServer(
			endpoints.RemoveRouteForPermissionEndpoint,
			decodeHttpRemoveRouteForPermissionRequest,
			encodeHttpResponse,
			httptransport.ServerErrorEncoder(errorEncoder),
		)))

	engine.POST("/user/addPermissionForRole",
		userservice.JwtMiddleware(keyFunc, stdjwt.SigningMethodHS256, userservice.UserClaimFactory),
		convertHttpHandlerToGinHandler(httptransport.NewServer(
			endpoints.AddPermissionForRoleEndpoint,
			decodeHttpAddPermissionForRoleRequest,
			encodeHttpResponse,
			httptransport.ServerErrorEncoder(errorEncoder),
		)))

}

func convertHttpHandlerToGinHandler(handler http.Handler) gin.HandlerFunc {
	return func(c *gin.Context) {
		handler.ServeHTTP(c.Writer, c.Request)
	}
}

func decodeHttpRegisterRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request pb.RegisterRequest
	defer r.Body.Close()
	err := json.NewDecoder(r.Body).Decode(&request)
	return &request, err
}

func decodeHttpLoginRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request pb.LoginRequest
	defer r.Body.Close()
	err := json.NewDecoder(r.Body).Decode(&request)
	return &request, err
}

func decodeHttpGetPermissionsRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request pb.GetPermissionsRequest
	defer r.Body.Close()
	err := json.NewDecoder(r.Body).Decode(&request)
	return &request, err
}

func decodeHttpCheckPermissionRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var request pb.CheckPermissionRequest
	defer r.Body.Close()
	err := json.NewDecoder(r.Body).Decode(&request)
	return &request, err
}

func decodeHttpAddRoutesRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request pb.AddRoutesRequest
	defer r.Body.Close()
	err := json.NewDecoder(r.Body).Decode(&request)
	return &request, err
}

func decodeHttpListRoutesRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request pb.ListRoutesRequest
	defer r.Body.Close()
	err := json.NewDecoder(r.Body).Decode(&request)
	return &request, err
}

func decodeHttpCreatePermissionRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request pb.CreatePermissionRequest
	defer r.Body.Close()
	err := json.NewDecoder(r.Body).Decode(&request)
	return &request, err
}

func decodeHttpUpdatePermissionRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request pb.UpdatePermissionRequest
	defer r.Body.Close()
	err := json.NewDecoder(r.Body).Decode(&request)
	return &request, err
}

func decodeHttpAddRouteForPermissionRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request pb.AddRouteForPermissionRequest
	defer r.Body.Close()
	err := json.NewDecoder(r.Body).Decode(&request)
	return &request, err
}

func decodeHttpRemoveRouteForPermissionRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request pb.RemoveRouteForPermissionRequest
	defer r.Body.Close()
	err := json.NewDecoder(r.Body).Decode(&request)
	return &request, err
}

func decodeHttpAddPermissionForRoleRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request pb.AddPermissionForRoleRequest
	defer r.Body.Close()
	err := json.NewDecoder(r.Body).Decode(&request)
	return &request, err
}

func encodeHttpResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	s, ok := response.(StatusError)
	if ok {
		err := s.GetErr()
		if err != nil {
			w.WriteHeader(int(err.GetCode()))
			return json.NewEncoder(w).Encode(gin.H{
				"err":  err.GetMessage(),
				"desc": err.GetDescription(),
			})
		}
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	return json.NewEncoder(w).Encode(response)
}

func errorEncoder(ctx context.Context, err error, w http.ResponseWriter) {
	code, msg := err2codeAndMessage(err)
	w.WriteHeader(code)
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	json.NewEncoder(w).Encode(gin.H{
		"err":  msg,
		"desc": "未知错误",
	})
}

func err2codeAndMessage(err error) (int, string) {
	if err == nil {
		return http.StatusOK, ""
	}
	switch e := err.(type) {
	case lb.RetryError:
		err = e.Final
	}
	return http.StatusInternalServerError, err.Error()
}

func keyFunc(token *jwt.Token) (interface{}, error) {
	return []byte(userservice.SignedKey), nil
}

func setUserInfoContext(ctx context.Context, r *http.Request) context.Context {
	username := r.Form.Get("username")
	id := r.Form.Get("id")
	md := metadata.New(map[string]string{
		"username": username,
		"id":       id,
	})
	return context.WithValue(ctx, "userInfo", md)
}
