package file

import (
	"bufio"
	"fmt"
	"go-lib/str"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

//覆盖写入文件
func RecoverWrite(path, content string) bool {
	file, err := os.OpenFile(path, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0666)
	if err != nil {
		return false
	}
	//及时关闭file句柄
	defer file.Close()
	//写入文件时，使用带缓存的 *Writer
	write := bufio.NewWriter(file)
	_, err = write.WriteString(content)
	if err != nil {
		return false
	}
	//Flush将缓存的文件真正写入到文件中
	err = write.Flush()
	return err == nil
}

//追加写入
func AppendWrite(path, content string) bool {
	file, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		return false
	}
	//及时关闭file句柄
	n, _ := file.Seek(0, io.SeekEnd)
	_, err = file.WriteAt([]byte(content), n)
	defer file.Close()
	return err == nil
}

//读取文件
func Read(path string) []byte {
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Println("[File Lib]", err.Error())
		return nil
	}
	return bytes
}

//读取文件信息
func ReadFile(path string) *os.File {
	f, err := os.OpenFile(path, os.O_RDONLY, 0600)
	if err != nil {
		fmt.Println("[File Lib]", err.Error())
		return nil
	}
	return f
}

//获取文件扩展名
func GetExt(path string) string {
	at := str.LastIndexOf(path, ".")
	if at > -1 {
		strLen := str.StringLen(path)
		ext := str.SubStr(path, at, strLen)
		return strings.ToLower(ext)
	}
	return ""
}

//图片类型判断
func IsImage(ext string) bool {
	if ext == ".jpg" || ext == ".bmp" || ext == ".jpeg" || ext == ".png" || ext == ".gif" {
		return true
	}
	return false
}

//创建目录
func MakeDir(path string) bool {
	flag, _ := ExistPath(path)
	if flag {
		return true
	}
	err := os.MkdirAll(path, os.ModePerm)
	return err == nil
}

//是否存在
func ExistPath(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}
