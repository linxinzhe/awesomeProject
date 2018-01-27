package stringutils

import (
	"strings"
	"fmt"
	"unicode/utf8"
	"encoding/hex"
	"log"
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

	n := 0
	for i, r := range "Hello, 世界" {
		fmt.Printf("%d\t%q\t%d\n", i, r, r)
		n++ //实际上遇到中文就跳过了3个字节
	}
	fmt.Println(n)

}

func DecodeByteHexString() {
	src := []byte("48656c6c6f20476f7068657221")

	dst := make([]byte, hex.DecodedLen(len(src)))
	n, err := hex.Decode(dst, src) //相当于两两，如"48"合并成一个hex真数字
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(dst[:n])  //hex真数字
	fmt.Printf("%s\n", dst[:n])
}
