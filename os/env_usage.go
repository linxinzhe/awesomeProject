package os

import (
	"os"
	"runtime"
)

// GetEnv gets environment var
func GetEnv() string {
	gopath := os.Getenv("GOPATH")
	return gopath
}

// Version gets your go version
func Version() string {
	version := runtime.Version()
	return version
}
