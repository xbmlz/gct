package gct

import (
	"archive/tar"
	"archive/zip"
	"bufio"
	"compress/gzip"
	"io"
	"os"
	"path"
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

// IsFileExists 判断文件是否存在
func (tf *TFile) IsFileExists(path string) bool {
	_, err := os.Lstat(path)
	return !os.IsNotExist(err)
}

// IsDirExists 判断目录是否存在
func (tf *TFile) IsDirExists(dirname string) bool {
	fi, err := os.Stat(dirname)
	return (err == nil || os.IsExist(err)) && fi.IsDir()
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

// Zip 压缩文件
func (tf *TFile) Zip(dest string, paths ...string) error {
	zfile, err := os.Create(dest)
	if err != nil {
		return err
	}
	defer zfile.Close()

	zipWriter := zip.NewWriter(zfile)
	defer zipWriter.Close()

	for _, src := range paths {
		// remove the trailing path sepeartor if it is a directory
		// src := strings.TrimSuffix(src, string(os.PathSeparator))

		err = filepath.Walk(src, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			// create local file header
			header, err := zip.FileInfoHeader(info)

			if err != nil {
				return err
			}
			// set relative path of file in zip archive
			header.Name, err = filepath.Rel(filepath.Dir(src), path)
			if err != nil {
				return err
			}
			if info.IsDir() {
				header.Name += string(os.PathSeparator)
			} else {
				// set compression method to deflate
				header.Method = zip.Deflate
			}
			// create writer for writing header
			headerWriter, err := zipWriter.CreateHeader(header)
			if err != nil {
				return err
			}
			if info.IsDir() {
				return nil
			}
			f, err := os.Open(path)
			if err != nil {
				return err
			}
			defer f.Close()
			_, err = io.Copy(headerWriter, f)
			return err
		})
		if err != nil {
			return err
		}
	}
	return nil
}

// Unzip 解压文件
func (tf *TFile) Unzip(src string, dest string) error {
	reader, err := zip.OpenReader(src)
	if err != nil {
		return err
	}
	defer reader.Close()

	for _, file := range reader.File {
		filePath := path.Join(dest, file.Name)
		if file.FileInfo().IsDir() {
			os.MkdirAll(filePath, os.ModePerm)
		} else {
			if err = os.MkdirAll(filepath.Dir(filePath), os.ModePerm); err != nil {
				return err
			}
			inFile, err := file.Open()
			if err != nil {
				return err
			}
			defer inFile.Close()

			outFile, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, file.Mode())
			if err != nil {
				return err
			}
			defer outFile.Close()

			_, err = io.Copy(outFile, inFile)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

// 压缩tar文件
func (tf *TFile) Gzip(dest string, paths ...string) error {
	tarfile, err := os.Create(dest)
	if err != nil {
		return err
	}
	defer tarfile.Close()

	gzipWriter := gzip.NewWriter(tarfile)
	defer gzipWriter.Close()

	tarWriter := tar.NewWriter(gzipWriter)
	defer tarWriter.Close()

	for _, src := range paths {
		// remove the trailing path sepeartor if it is a directory
		src := strings.TrimSuffix(src, string(os.PathSeparator))

		err = filepath.Walk(src, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			// create local file header
			header, err := tar.FileInfoHeader(info, "")
			if err != nil {
				return err
			}
			// set relative path of file in zip archive
			header.Name, err = filepath.Rel(filepath.Dir(src), path)
			if err != nil {
				return err
			}
			if info.IsDir() {
				header.Name += string(os.PathSeparator)
			}
			// create writer for writing header
			err = tarWriter.WriteHeader(header)
			if err != nil {
				return err
			}
			if info.IsDir() {
				return nil
			}
			f, err := os.Open(path)
			if err != nil {
				return err
			}
			defer f.Close()
			_, err = io.Copy(tarWriter, f)
			return err
		})
		if err != nil {
			return err
		}
	}
	return nil
}

// 解压tar文件
func (tf *TFile) UnGzip(src, dest string) error {
	fr, err := os.Open(src)
	if err != nil {
		return err
	}
	defer fr.Close()

	gr, err := gzip.NewReader(fr)
	if err != nil {
		return err
	}
	defer gr.Close()

	tr := tar.NewReader(gr)

	for {
		hdr, err := tr.Next()
		if err != nil {
			if err == io.EOF {
				break
			}
			return err
		}
		destName := filepath.Join(dest, hdr.Name)
		if hdr.Typeflag == tar.TypeDir {
			if b := FileUtils.IsDirExists(destName); !b {
				if err := os.MkdirAll(destName, 0775); err != nil {
					return err
				}
			}
		} else {
			file, err := os.OpenFile(destName, os.O_CREATE|os.O_RDWR, os.FileMode(hdr.Mode))
			if err != nil {
				return err
			}
			_, err = io.Copy(file, tr)
			if err != nil {
				return err
			}
			file.Close()
		}
	}
	return nil
}

func IsDirExists(destName string) {
	panic("unimplemented")
}
