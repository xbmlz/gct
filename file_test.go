package gct

import "testing"

func TestExist(t *testing.T) {
	path := "./file.go"
	if !FileUtils.Exist(path) {
		t.Error("file.go must exist")
		return
	}
}
