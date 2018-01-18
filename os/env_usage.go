package os

import (
	"os"
	"runtime"
)

func GetEnv() string {
	gopath := os.Getenv("GOPATH")
	return gopath
}

func Version() string {
	version := runtime.Version()
	return version
}
