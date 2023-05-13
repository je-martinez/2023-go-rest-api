package router

import (
	auth_handlers "github.com/je-martinez/2023-go-rest-api/api/v1/handlers/auth"
	post_handlers "github.com/je-martinez/2023-go-rest-api/api/v1/handlers/post"
	server_handlers "github.com/je-martinez/2023-go-rest-api/api/v1/handlers/server"
	user_handlers "github.com/je-martinez/2023-go-rest-api/api/v1/handlers/user"
	"github.com/je-martinez/2023-go-rest-api/api/v1/middleware"
	"github.com/je-martinez/2023-go-rest-api/pkg/constants"
	"github.com/je-martinez/2023-go-rest-api/pkg/logger"
	router_types "github.com/je-martinez/2023-go-rest-api/pkg/types/router"

	"github.com/gin-gonic/gin"
)

func New(address string, logger *logger.ApiLogger, props *router_types.RouterHandlerProps) *RouterApiInstance {
	return &RouterApiInstance{
		props:   props,
		logger:  logger,
		address: address,
		gin:     gin.New(),
	}
}

type RouterApiInstance struct {
	props   *router_types.RouterHandlerProps
	logger  *logger.ApiLogger
	address string
	gin     *gin.Engine
}

func (r *RouterApiInstance) RegisterRoutes() {
	publicPath := "/api/v1/public"
	protectedRelativePath := "/api/v1"

	//Router Groups
	GinPublic := r.gin.Group(publicPath)
	{
		//Auth
		GinPublic.POST(constants.Login, auth_handlers.Login(r.props))
		GinPublic.POST(constants.RegisterUser, auth_handlers.RegisterUser(r.props))

		//Server
		GinPublic.GET(constants.Health, server_handlers.Health(r.props))
		GinPublic.GET(constants.Metrics, server_handlers.PrometheusHandler(r.props))
	}

	GinProtected := r.gin.Group(protectedRelativePath, middleware.AuthMiddleware(r.props))
	{
		//User
		GinProtected.GET(constants.Me, user_handlers.Me(r.props))
		GinProtected.PUT(constants.UpdateUser, user_handlers.UpdateUser(r.props))

		//Post
		GinProtected.POST(constants.CreatePost, post_handlers.CreatePost(r.props))

		//Server
		GinProtected.GET(constants.HealthAuth, server_handlers.Health(r.props))
	}
}

func (r *RouterApiInstance) Start() {
	err := r.gin.Run(r.address)
	if err != nil {
		r.logger.Fatalf(constants.API_RUNNING, err)
	}
}
