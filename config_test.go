package codebase

import (
	"testing"
)

func Test_HelloWorld(t *testing.T) {
	hello := HelloWorld()
	if hello != "Hello, World!" {
		t.Errorf("not hello world")
	}
}
