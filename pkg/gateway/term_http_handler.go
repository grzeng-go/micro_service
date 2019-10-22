package gateway

import (
	"userService/pkg/pb"

	"github.com/gin-gonic/gin"
	httptransport "github.com/go-kit/kit/transport/http"
)

func RegisterTermHandler(engine *gin.Engine, endpoints *TermEndpoints) {
	group := engine.Group("/term")

	group.POST("/listTermInfo", convertHttpHandlerToGinHandler(httptransport.NewServer(
		endpoints.ListTermInfoEndpoint,
		decodeHttpRequest(&pb.ListTermInfoRequest{}),
		encodeHttpResponse,
		httptransport.ServerErrorEncoder(errorEncoder),
	)))

	group.POST("/saveTerm", convertHttpHandlerToGinHandler(httptransport.NewServer(
		endpoints.SaveTermEndpoint,
		decodeHttpRequest(&pb.SaveTermRequest{}),
		encodeHttpResponse,
		httptransport.ServerErrorEncoder(errorEncoder),
	)))

	group.POST("/listTermRisk", convertHttpHandlerToGinHandler(httptransport.NewServer(
		endpoints.ListTermRiskEndpoint,
		decodeHttpRequest(&pb.ListTermRiskRequest{}),
		encodeHttpResponse,
		httptransport.ServerErrorEncoder(errorEncoder),
	)))

	group.POST("/listTermActivationState", convertHttpHandlerToGinHandler(httptransport.NewServer(
		endpoints.ListTermActivationStateEndpoint,
		decodeHttpRequest(&pb.ListTermActivationStateRequest{}),
		encodeHttpResponse,
		httptransport.ServerErrorEncoder(errorEncoder),
	)))

	group.POST("/updateTermInfo", convertHttpHandlerToGinHandler(httptransport.NewServer(
		endpoints.UpdateTermInfoEndpoint,
		decodeHttpRequest(&pb.UpdateTermInfoRequest{}),
		encodeHttpResponse,
		httptransport.ServerErrorEncoder(errorEncoder),
	)))
}
