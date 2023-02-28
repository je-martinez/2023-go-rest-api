package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func Health(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"ok":     true,
		"status": "Everything is okay!",
	})
}

func PrometheusHandler() gin.HandlerFunc {
	h := promhttp.Handler()

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}
