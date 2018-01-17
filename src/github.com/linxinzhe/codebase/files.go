package main

import (
	"os"
	"path/filepath"
	"log"
)

func main() {
	path := filepath.Join("./src/github.com/linxinzhe/codebase", "files.go")
	log.Println(path)

	_, err := os.Stat(path)

	if os.IsNotExist(err) {
		log.Fatal("file not exist")
	} else {
		log.Println("file exist")
	}

}
