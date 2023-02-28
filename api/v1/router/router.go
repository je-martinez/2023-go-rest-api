package router

import (
	handler "main/api/v1/handlers"
	routes "main/constants/routes"

	"github.com/gin-gonic/gin"
)

func Start(r *gin.Engine) {
	//Internal Handlers
	r.GET(routes.Health, handler.Health)
	r.GET(routes.Metrics, handler.PrometheusHandler())
}
