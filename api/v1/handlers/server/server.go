package server_handlers

import (
	"github.com/je-martinez/2023-go-rest-api/pkg/constants"
	router_types "github.com/je-martinez/2023-go-rest-api/pkg/types/router"
	"github.com/je-martinez/2023-go-rest-api/pkg/utils"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func Health(props *router_types.RouterHandlerProps) gin.HandlerFunc {
	return gin.HandlerFunc(func(c *gin.Context) {
		utils.GinApiResponse(c, 200, constants.MSG_HEALTH, nil, nil)
	})
}

func PrometheusHandler(props *router_types.RouterHandlerProps) gin.HandlerFunc {
	h := promhttp.Handler()
	return gin.HandlerFunc(func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	})
}
