package handlers

import (
	"net/http"

	"main/pkg/utils"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func Health(c *gin.Context) {
	c.JSON(http.StatusOK, utils.SuccessApiResponse("Everything is okay", 200))
}

func PrometheusHandler() gin.HandlerFunc {
	h := promhttp.Handler()

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}
