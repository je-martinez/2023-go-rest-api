package utils

import "main/pkg/types"

// Get config path for local or docker
func GetConfigPath(configPath string) string {
	if configPath == "" {
		return "./config/config"
	}
	return configPath
}

func SuccessApiResponse(data interface{}, statusCode int) (response *types.ApiResponse) {
	return &types.ApiResponse{
		OK:         true,
		StatusCode: statusCode,
		Data:       data,
		Errors:     []string{},
	}
}
