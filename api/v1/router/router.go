package router

import (
	auth_handlers "main/api/v1/handlers/auth"
	user_handlers "main/api/v1/handlers/user"
	"main/api/v1/middleware"

	sv_handlers "main/api/v1/handlers/server"
	routes "main/pkg/constants"

	"github.com/gin-gonic/gin"
)

func Start(r *gin.Engine) {
	//Internal Handlers
	r.GET(routes.Health, sv_handlers.Health)
	r.GET(routes.Metrics, sv_handlers.PrometheusHandler())

	//Auth
	r.POST(routes.Login, auth_handlers.Login)
	r.POST(routes.RegisterUser, auth_handlers.RegisterUser)

	//User
	r.PUT(routes.UpdateUser, middleware.AuthMiddleware(), user_handlers.UpdateUser)

}
