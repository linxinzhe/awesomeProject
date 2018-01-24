package stringutils

import (
	"regexp"
)

// FindStringExample find the string by using regex
func FindStringExample() string {
	str := "abcd1234"
	r := regexp.MustCompile(`[0-9]+`)
	result := r.FindString(str)
	return result
}
