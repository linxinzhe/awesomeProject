package stringutils

import (
	"strings"
)

func contains(text string, str string) bool {
	return strings.Contains(text, str)
}