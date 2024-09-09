package utils

import (
	"strings"
)

func FormatErrors(s string) string {
	return strings.ReplaceAll(s, "\n", ", ")
}
