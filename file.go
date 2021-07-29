package gct

import (
	"bufio"
	"os"
)

// AppendString 将String写入文件，追加模式
func (tf *TFile) AppendString(content string, path string) (*os.File, error) {
	file, err := os.OpenFile(path, os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	write := bufio.NewWriter(file)
	write.WriteString(content)
	write.Flush()
	return file, nil
}

// Exist 判断文件是否存在
func (tf *TFile) Exist(path string) bool {
	_, err := os.Lstat(path)
	return !os.IsNotExist(err)
}
