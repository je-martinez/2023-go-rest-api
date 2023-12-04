package constants

const (
	//Server
	Health     = "/health"
	HealthAuth = "/health-with-auth"
	Metrics    = "/metrics"
	//Auth
	Login        = "/auth/login"
	RegisterUser = "/auth/register"
	//User
	Me         = "/user/me"
	UpdateUser = "/user/update"
	//Post
	CreatePost = "/post/create"
	DeletePost = "/post/:post_id"
	UpdatePost = "/post/:post_id"
	//Reaction
	CreatePostReaction = "/reaction/post/:post_id/:reaction_type"
	DeletePostReaction = "/reaction/post/:post_id/:reaction_type"
)
