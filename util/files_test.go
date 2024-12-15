package util

import "testing"

func TestFileList(t *testing.T) {
	val, err := FileList("./exmple/")
	if err != nil {
		t.Log(val)
	}
}
