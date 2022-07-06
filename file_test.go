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

func TestRemoveSuffix(t *testing.T) {
	var tests = []struct {
		param    string
		expected string
	}{
		{"file.go", "file"},
		{"file.Go", "file"},
		{"file.tar.gz", "file.tar"},
	}
	for _, test := range tests {
		actual := FileUtils.RemoveSuffix(test.param)
		if actual != test.expected {
			t.Errorf("%s must be %s", test.param, test.expected)
		}
	}
}

func TestRemovePrefix(t *testing.T) {
	var tests = []struct {
		param    string
		expected string
	}{
		{"file.go", "go"},
		{"file.Go", "Go"},
		{"file.tar.gz", "gz"},
	}
	for _, test := range tests {
		actual := FileUtils.RemovePrefix(test.param)
		if actual != test.expected {
			t.Errorf("%s must be %s", test.param, test.expected)
		}
	}
}

func TestZip(t *testing.T) {
	err := FileUtils.Zip("./testdata/archive.zip", "./testdata/csv", "./testdata/file.txt")
	if err != nil {
		t.Error("zip error")
		return
	}
}

func TestUnzip(t *testing.T) {
	err := FileUtils.Unzip("./testdata/archive.zip", "./testdata/unzip")
	if err != nil {
		t.Error("unzip error")
		return
	}
}
