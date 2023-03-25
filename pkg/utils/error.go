package utils

import (
	"strings"
)

func EntityNotFound(err error) bool {
	if err == nil {
		return false
	}
	return strings.Contains("record not found", err.Error())
}
