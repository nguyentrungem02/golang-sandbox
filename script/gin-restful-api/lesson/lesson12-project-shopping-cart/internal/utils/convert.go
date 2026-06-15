package utils

import (
	"regexp"
	"strings"
)

var (
	matchFirstCap = regexp.MustCompile("(.)([A-Z][a-z]+)")
	matchAllCap   = regexp.MustCompile("([a-z0-9])([A-Z])")
)

func PascalToSnake(s string) string {
	snake := matchFirstCap.ReplaceAllString(s, "${1}_${2}")
	snake = matchAllCap.ReplaceAllString(snake, "${1}_${2}")

	return strings.ToLower(snake)
}

func NormalizeString(text string) string {
	return strings.ToLower(strings.TrimSpace(text))
}

func ConvertToInt32Pointer(val int) *int32 {
	if val == 0 {
		return nil
	}

	v := int32(val)
	return &v
}

func CapitalizeFirst(s string) string {
	if len(s) == 0 {
		return s
	}

	return strings.ToUpper(s[0:1]) + s[1:]
}
