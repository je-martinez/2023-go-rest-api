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

func GinApiResponse(c *gin.Context, statusCode int, data any, errors any) {
	c.JSON(http.StatusOK, apiResponse(200, data, errors))
}

func apiResponse(statusCode int, data any, errors any) (response *types.ApiResponse) {

	if errors == nil {
		errors = []string{}
	}

	return &types.ApiResponse{
		OK:         isResponseSuccess(statusCode),
		StatusCode: statusCode,
		Data:       data,
		Errors:     errors,
	}
}

func isResponseSuccess(statusCode int) bool {
	return statusCode >= 200 && statusCode <= 299
}
