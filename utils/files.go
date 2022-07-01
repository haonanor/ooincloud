package utils

import (
	"os"
	"strings"

	uuid "github.com/satori/go.uuid"
)

func DirExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}

	if os.IsNotExist(err) {
		return false, nil
	}

	return false, err

}

func OpenFile(filename string) (*os.File, error) {
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		return os.Create(filename) //创建文件
	}
	return os.OpenFile(filename, os.O_APPEND, 0666) //打开文件
}

func CreateFileKey() string {
	u := uuid.NewV4().String()
	return u
}

// 获取文件后缀
func GetFileExt(filename string) string {
	middleFileData := strings.Split(filename, ".")
	exp := middleFileData[len(middleFileData)-1]
	return exp
}
