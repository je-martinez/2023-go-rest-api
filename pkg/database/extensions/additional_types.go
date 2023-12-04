package db_extensions

import "fmt"

type SignInProviderType string

const (
	EMAIL  SignInProviderType = "email"
	GOOGLE SignInProviderType = "google"
	APPLE  SignInProviderType = "apple"
)

const (
	LOVE ReactionType = "love"
)

func GetSupportedReactions() []ReactionType {
	return []ReactionType{LOVE}
}

type ReactionType string

func (r ReactionType) ToString() string {
	return fmt.Sprint(r)
}
