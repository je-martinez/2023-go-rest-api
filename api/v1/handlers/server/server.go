package server_handlers

import (
	"main/pkg/constants"
	"main/pkg/utils"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func Health(c *gin.Context) {
	utils.GinApiResponse(c, 200, constants.MSG_HEALTH, nil, nil)
}

func PrometheusHandler() gin.HandlerFunc {
	h := promhttp.Handler()

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}
