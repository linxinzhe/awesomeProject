package io

import (
	"fmt"
	"os"
)

type Data struct {
}

func (self Data) String() string {
	return "data"
}

func FmtUsage() {
	fmt.Printf("hello world\n")

	fmt.Println("hello world")

	fmt.Printf("num %d\n", 5)

	str := fmt.Sprintf("float %f", 3.14159)
	fmt.Print(str)

	fmt.Fprintln(os.Stdout, "A\n")

	fmt.Printf("%v\n", Data{})

}
