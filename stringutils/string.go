package stringutils

import (
	"strings"
	"fmt"
	"unicode/utf8"
)

// Contains check if str is in text
func Contains(text string, str string) bool {
	return strings.Contains(text, str)
}

func stringToByte() {
	a := "啊bcd" // utf-8，一个中文顶三个byte，一个英文一个byte
	fmt.Println([]byte(a))
}

func lengthOfUTF8() {
	s := "Hello, 世界"
	fmt.Println(len(s))                    // "13"
	fmt.Println(utf8.RuneCountInString(s)) // "9"
}
