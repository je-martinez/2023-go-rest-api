package router

import (
	sv_handlers "main/api/v1/handlers/server"
	routes "main/pkg/constants"

	"github.com/gin-gonic/gin"
)

func Start(r *gin.Engine) {
	//Internal Handlers
	r.GET(routes.Health, sv_handlers.Health)
	r.GET(routes.Metrics, sv_handlers.PrometheusHandler())

	//

}
