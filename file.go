package gct

import "os"

// Exist 判断文件是否存在
func (f *File) Exist(path string) bool {
	_, err := os.Lstat(path)
	return !os.IsNotExist(err)
}
