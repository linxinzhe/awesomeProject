package file

import (
	"testing"
)

func Test_Exist(t *testing.T) {
	exist := Exist(".","file_exist.go")
	if !exist {
		t.Errorf("Expected exist")
	}
}
