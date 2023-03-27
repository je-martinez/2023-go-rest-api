package constants

// General Errors
const (
	ERR_BIND_JSON      = "Error binding JSON"
	ERR_INVALID_JSON   = "Error with the provided JSON"
	ERR_CURRENT_USER   = "Error related to retrieve logged user"
	ERR_GENERATE_TOKEN = "Error trying to generate a new access token"
)

// Entity Errors
const (
	ERR_ENTITY_NOT_FOUND    = "%s not found in the database"
	ERR_ENTITY_NOT_FOUND_ID = "%s not found with id: %s"
	ERR_CREATE_ENTITY       = "Error trying to create <%s>"
)

// User Entity Errros
const (
	ERR_USERNAME_PASSWORD_INVALID = "Username or password are invalid"
)
