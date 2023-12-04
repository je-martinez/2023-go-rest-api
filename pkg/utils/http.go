package utils

import (
	"net/http"

	types "github.com/je-martinez/2023-go-rest-api/pkg/types/http"

	"github.com/gin-gonic/gin"
)

func GinApiResponse(c *gin.Context, statusCode int, message *string, data any, err any) {
	c.JSON(http.StatusOK, apiResponse(statusCode, message, data, err))
}

func apiResponse(statusCode int, message *string, data any, err any) (response *types.ApiResponse) {

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
