package gct

import "testing"

func TestAppendString(t *testing.T) {
	path := "./testdata/file.txt"
	_, err := FileUtils.AppendString("test append string", path)
	if err != nil {
		t.Error("file append error")
		return
	}
}

func TestExist(t *testing.T) {
	path := "./file.go"
	if !FileUtils.Exist(path) {
		t.Error("file.go must exist")
		return
	}
}
