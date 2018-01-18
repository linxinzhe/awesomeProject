package os

import (
	"os"
)

func GetEnv() string {
	gopath := os.Getenv("GOPATH")
	return gopath
}
