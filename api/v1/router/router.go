package router

import (
	auth_handlers "github.com/je-martinez/2023-go-rest-api/api/v1/handlers/auth"
	post_handlers "github.com/je-martinez/2023-go-rest-api/api/v1/handlers/post"
	user_handlers "github.com/je-martinez/2023-go-rest-api/api/v1/handlers/user"
	"github.com/je-martinez/2023-go-rest-api/api/v1/middleware"

	sv_handlers "github.com/je-martinez/2023-go-rest-api/api/v1/handlers/server"
	routes "github.com/je-martinez/2023-go-rest-api/pkg/constants"

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

		//Post
		GinProtected.POST(routes.CreatePost, post_handlers.CreatePost)

		//Server
		GinProtected.GET(routes.HealthAuth, sv_handlers.Health)
	}
}
