package utils

import (
	"strings"
)

func CleanErrors(s string) string {
	return strings.ReplaceAll(s, "\n", ", ")
}
