package router

import (
	routes "main/constants/routes"
	handler "main/handlers"

	"github.com/gin-gonic/gin"
)

func Start(r *gin.Engine) {
	//Internal Handlers
	r.GET(routes.Health, handler.Health)
}
