package gct

import (
	"bufio"
	"os"
	"path/filepath"
	"strings"
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

// RemoveSuffix 删除文件后缀
func (tf *TFile) RemoveSuffix(path string) string {
	suffix := filepath.Ext(path)
	if suffix != "" {
		return strings.Replace(path, suffix, "", -1)
	}
	return ""
}

// RemovePrefix 删除文件前缀
func (tf *TFile) RemovePrefix(path string) string {
	prefix := filepath.Ext(path)
	if prefix != "" {
		return prefix[1:]
	}
	return ""
}
