package utils

import (
	"fmt"
	"main/config"
	"main/pkg/database/entities"
	"main/pkg/types"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
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

func ExtractToken(c *gin.Context) string {
	token := c.Query("token")
	if token != "" {
		return token
	}
	bearerToken := c.Request.Header.Get("Authorization")
	if len(strings.Split(bearerToken, " ")) == 2 {
		return strings.Split(bearerToken, " ")[1]
	}
	return ""
}

func TokenValid(c *gin.Context) error {
	SECRET_KEY := config.AppConfig.Server.JwtSecretKey
	tokenString := ExtractToken(c)
	_, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(SECRET_KEY), nil
	})
	if err != nil {
		return err
	}
	return nil
}

func ExtractUserFromToken(c *gin.Context) (*types.CurrentUser, error) {
	SECRET_KEY := config.AppConfig.Server.JwtSecretKey
	tokenString := ExtractToken(c)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(SECRET_KEY), nil
	})
	if err != nil {
		return nil, err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		user := &types.CurrentUser{}
		user.FromClaims(claims)
		return user, nil
	}
	return nil, nil
}
