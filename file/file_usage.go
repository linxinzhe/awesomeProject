package file

import (
	"os"
	"path/filepath"
	"log"
)

// check if the file is in dir
func Exist(dir string, filename string) bool {
	path := filepath.Join(dir, filename)
	log.Println("Path", path)

	_, err := os.Stat(path)
	log.Println("err", err)

	return !os.IsNotExist(err)
}

// your file absolute path
func Abs(path string, filename string) string {
	var fileAbsPath, _ = filepath.Abs(filepath.Join(path, filename))
	return fileAbsPath
}

