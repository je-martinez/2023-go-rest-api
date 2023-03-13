package utils

import (
	"main/pkg/types"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Get config path for local or docker
func GetConfigPath(configPath string) string {
	if configPath == "" {
		return "./config/config"
	}
	return configPath
}

func GinApiResponse(c *gin.Context, statusCode int, message string, data any, err any) {
	c.JSON(http.StatusOK, apiResponse(200, message, data, err))
}

func apiResponse(statusCode int, message string, data any, err any) (response *types.ApiResponse) {

	return &types.ApiResponse{
		OK:         isResponseSuccess(statusCode),
		StatusCode: statusCode,
		Message:    message,
		Data:       data,
		Error:      err,
	}
}

func isResponseSuccess(statusCode int) bool {
	return statusCode >= 200 && statusCode <= 299
}
