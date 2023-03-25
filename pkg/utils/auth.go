package utils

import (
	"main/config"
	"main/pkg/database/entities"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateToken(user entities.User) (token string, err error) {
	claimsMap := jwt.MapClaims{
		"user_id":          user.UserID,
		"username":         user.Username,
		"fullname":         user.Fullname,
		"sign_in_provider": user.SignInProvider,
		"email":            user.Email,
		"exp":              time.Now().Add(time.Hour * 24).Unix(),
	}

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claimsMap)
	SECRET_KEY := config.AppConfig.Server.JwtSecretKey
	tokenString, err := jwtToken.SignedString([]byte(SECRET_KEY))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
