package util

import (
	"strings"

	"github.com/google/uuid"
)

func NewUUID() string {
	return uuid.New().String()
}

func WhitelistSanitize(str string, whitelist map[string]string) string {
	for k, v := range whitelist {
		str = strings.ReplaceAll(str, k, v)
	}
	return str
}
