package file

import (
	"os"
	"path/filepath"
	"log"
)

func Exist(dir string, filename string) bool {
	path := filepath.Join(dir, filename)
	log.Println("Path", path)

	_, err := os.Stat(path)
	log.Println("err", err)

	return !os.IsNotExist(err)
}
