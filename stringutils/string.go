package stringutils

import (
	"strings"
)

// check if str is in text
func contains(text string, str string) bool {
	return strings.Contains(text, str)
}