package db_extensions

import "database/sql/driver"

type SignInProviderType string

const (
	EMAIL  SignInProviderType = "email"
	GOOGLE SignInProviderType = "google"
	APPLE  SignInProviderType = "apple"
)

func (p *SignInProviderType) Scan(value interface{}) error {
	*p = SignInProviderType(value.([]byte))
	return nil
}

func (p SignInProviderType) Value() (driver.Value, error) {
	return string(p), nil
}

const (
	LIKE = "like"
)

type ReactionType string

func (p *ReactionType) Scan(value interface{}) error {
	*p = ReactionType(value.([]byte))
	return nil
}

func (p ReactionType) Value() (driver.Value, error) {
	return string(p), nil
}
