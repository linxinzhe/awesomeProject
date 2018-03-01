package io

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

//http://www.cnblogs.com/golove/p/3282667.html
func BufferUsage() {

	strReader := strings.NewReader("hello world")

	bufReader := bufio.NewReader(strReader)

	// 只读拾取
	data, _ := bufReader.Peek(5)

	fmt.Println(data, bufReader.Buffered())

	// 读出hello ， 查看剩余
	str, _ := bufReader.ReadString(' ')

	// 取出
	fmt.Println(str, bufReader.Buffered())

	// 写入
	w := bufio.NewWriter(os.Stdout)

	fmt.Fprint(w, "Hello, ")
	fmt.Fprint(w, "world!")
	w.Flush() // Don't forget to flush!

}
