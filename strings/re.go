package strings

import (
	"regexp"
)

func FindStringExample() string {
	str := "abcd1234"
	r := regexp.MustCompile(`[0-9]+`)
	result := r.FindString(str)
	return result
}
