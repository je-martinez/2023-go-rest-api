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

	publicPath := "/api/v1/public"
	protectedRelativePath := "/api/v1"

	//Router Groups
	GinPublic := r.Group(publicPath)
	{
		//Auth
		GinPublic.POST(routes.Login, auth_handlers.Login)
		GinPublic.POST(routes.RegisterUser, auth_handlers.RegisterUser)

		//Server
		GinPublic.GET(routes.HealthAuth, sv_handlers.Health)
		GinPublic.GET(routes.Metrics, sv_handlers.PrometheusHandler())
	}

	GinProtected := r.Group(protectedRelativePath, middleware.AuthMiddleware())
	{
		//User
		GinProtected.GET(routes.Me, user_handlers.Me)
		GinProtected.PUT(routes.UpdateUser, user_handlers.UpdateUser)

		//Server
		GinProtected.GET(routes.HealthAuth, sv_handlers.Health)
	}
}
