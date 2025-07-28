package utils

import "strings"

func ParseCallback(data string) string {
	clean := strings.TrimPrefix(data, "\f")
	parts := strings.Split(clean, "|")

	return parts[0]
}
