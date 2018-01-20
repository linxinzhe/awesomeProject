package os

import (
	"os"
	"runtime"
)

// get environment var
func GetEnv() string {
	gopath := os.Getenv("GOPATH")
	return gopath
}

// get your go version
func Version() string {
	version := runtime.Version()
	return version
}
