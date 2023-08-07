package constants

// General Errors
const (
	ERR_BIND_JSON      = "Error binding JSON"
	ERR_BIND_MULTIPART = "Error binding multipart/form data"
	ERR_INVALID_JSON   = "Error with the provided JSON"
	ERR_CURRENT_USER   = "Error related to retrieve logged user"
	ERR_GENERATE_TOKEN = "Error trying to generate a new access token"
	ERR_GENERATE_HASH  = "Error trying to generate a new %s hash"
)

// Entity Errors
const (
	ERR_ENTITY_NOT_FOUND    = "%s not found in the database"
	ERR_ENTITY_NOT_FOUND_ID = "%s not found with id: %s"
	ERR_FIND_ENTITY         = "Unable to trying to retrive entity %s"
	ERR_CREATE_ENTITY       = "Unable to create entity %s"
	ERR_UPDATE_ENTITY       = "Unable to update entity %s"
)

// User Entity Errros
const (
	ERR_USERNAME_PASSWORD_INVALID = "Username or password are invalid"
	ERR_OLD_PASSWORD_MISMATCH     = "Old password provide is invalid"
)
