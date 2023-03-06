package utils

// Get config path for local or docker
func GetConfigPath(configPath string) string {
	if configPath == "" {
		return "./config/config"
	}
	return configPath
}
