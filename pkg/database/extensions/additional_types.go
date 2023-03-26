package db_extensions

type SignInProviderType string

const (
	EMAIL  SignInProviderType = "email"
	GOOGLE SignInProviderType = "google"
	APPLE  SignInProviderType = "apple"
)

const (
	LIKE = "like"
)

type ReactionType string
