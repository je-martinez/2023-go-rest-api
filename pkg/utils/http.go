package utils

import (
	"main/pkg/types"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GinApiResponse(c *gin.Context, statusCode int, message string, data any, err any) {
	c.JSON(http.StatusOK, apiResponse(statusCode, message, data, err))
}

func apiResponse(statusCode int, message string, data any, err any) (response *types.ApiResponse) {

	return &types.ApiResponse{
		OK:         isResponseSuccess(statusCode),
		StatusCode: statusCode,
		Message:    message,
		Data:       data,
		Errors:     err,
	}
}

func isResponseSuccess(statusCode int) bool {
	return statusCode >= 200 && statusCode <= 299
}
