package util

import (
	"bytes"
	"os"
)

//default buff size
const defaultBuffSize = 1000

//读取文件
func ReadFromFile(fileName string, content *bytes.Buffer) error {
	f, err := os.Open(fileName)
	if err != nil {
		return err
	}
	defer f.Close()
	_, err = content.ReadFrom(f)
	return err
}

//从本地读取JSON
func GetLocalJSONBytes(path string) (*bytes.Buffer, error) {
	content := bytes.NewBuffer(make([]byte, 0, defaultBuffSize))
	err := ReadFromFile(path, content)
	return content, err
}
