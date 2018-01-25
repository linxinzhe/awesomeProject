package stringutils

import (
	"strings"
	"fmt"
)

// Contains check if str is in text
func Contains(text string, str string) bool {
	return strings.Contains(text, str)
}

func StringToByte() {
	a := "啊bcd" // utf-8，一个中文顶三个byte，一个英文一个byte
	fmt.Println([]byte(a))
}
