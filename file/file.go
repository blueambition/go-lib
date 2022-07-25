package file

import (
	"beego-base/functions/str"
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

// IsHidden checks whether the file specified by the given path is hidden.
func IsHidden(path string) bool {
	path = filepath.Base(path)
	if len(path) < 1 {
		return false
	}
	return "." == path[:1]
}

// GetFileSize get the length in bytes of file of the specified path.
func GetSize(path string) int64 {
	fi, err := os.Stat(path)
	if err != nil {
		return -1
	}
	return fi.Size()
}

// IsExist determines whether the file spcified by the given path is exists.
func IsExist(path string) bool {
	_, err := os.Stat(path)
	return err == nil || os.IsExist(err)
}

// IsDir determines whether the specified path is a directory.
func IsDir(path string) bool {
	fio, err := os.Lstat(path)
	if os.IsNotExist(err) {
		return false
	}
	if err != nil {
		return false
	}
	return fio.IsDir()
}

//获取文件扩展名(.txt)
func GetExt(filePath string) string {
	at := strings.LastIndex(filePath, ".")
	if at > -1 {
		strLen := str.Len(filePath)
		ext := str.SubStr(filePath, at, strLen)
		return strings.ToLower(ext)
	}
	return ""
}

func IsImage(filePath string) bool {
	ext := GetExt(filePath)
	if ext == ".jpg" || ext == ".bmp" || ext == ".jpeg" || ext == ".png" || ext == ".gif" {
		return true
	}
	return false
}

//读取文件
func ReadFile(filePath string) string {
	fileBytes, err := ioutil.ReadFile(filePath)
	if err != nil {
		return ""
	}
	return string(fileBytes)
}

//追加写入文件
func AppendWrite(filePath string, content string) error {
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		return err
	}
	defer file.Close()
	write := bufio.NewWriter(file)
	_, err = write.WriteString(content)
	if err != nil {
		return err
	}
	err = write.Flush()
	return err
}

//覆盖写入
func RecoverWrite(filePath string, content string) error {
	f, err := os.OpenFile(filePath, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0644)
	if err != nil {
		return err
	}
	n, _ := f.Seek(0, io.SeekEnd)
	_, err = f.WriteAt([]byte(content), n)
	fmt.Println("write succeed!")
	defer f.Close()
	return err
}
