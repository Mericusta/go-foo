package utility

import (
	"bufio"
	"os"
)

// ReadFileLineOneByOne 逐行读取文件内容，执行函数返回 true 则继续读取，返回 false 则结束读取
func ReadFileLineOneByOne(filename string, f func(string) bool) error {
	file, openError := os.Open(filename)
	if openError != nil {
		return openError
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		if !f(scanner.Text()) {
			break
		}
	}

	if scanner.Err() != nil {
		return scanner.Err()
	}

	return nil
}

// IsExist 检查文件或文件夹是否存在
func IsExist(path string) bool {
	_, err := os.Stat(path)
	return err == nil || os.IsExist(err)
}
