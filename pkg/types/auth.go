package types

import (
	"fmt"

	"github.com/golang-jwt/jwt/v5"
)

type CurrentUser struct {
	UserID         string
	Username       string
	Fullname       string
	SignInProvider string
	Email          string
}

func (u *CurrentUser) FromClaims(claims jwt.MapClaims) {
	u.UserID = fmt.Sprintf("%s", claims["user_id"])
	u.Username = fmt.Sprintf("%s", claims["username"])
	u.Fullname = fmt.Sprintf("%s", claims["fullname"])
	u.SignInProvider = fmt.Sprintf("%s", claims["sign_in_provider"])
	u.Email = fmt.Sprintf("%s", claims["email"])
}
