package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func ReadFile(filename string) (string, error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("ioutil.ReadFile error", err)
		return "", err
	}
	return string(data), nil
}

// 覆盖写
// 文件不存在则直接创建
func WriteFile(filename, content string) error {
	err := ioutil.WriteFile(filename, []byte(content), 0777)
	if err != nil {
		fmt.Println("ioutil.WriteFile error", err)
		return err
	}
	return nil
}

// 追加写
// 文件不存在则直接创建
func AppendFile(filename, content string) error {
	f, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("os.OpenFile error", err)
		return err
	}
	defer f.Close()

	if _, err = f.WriteString(content); err != nil {
		fmt.Println("file WriteString error", err)
		return err
	}
	return nil
}

var (
	filename = "mydata.txt"
)

func main() {
	WriteFile(filename, "Hello,World!你好，世界。\n")
	WriteFile(filename, "Hello,World!你好，世界。\n")
	content, _ := ReadFile(filename)
	fmt.Printf("file content: %s\n", content)

	AppendFile(filename, "Hello,World!你好，世界。\n")
	AppendFile(filename, "Hello,World!你好，世界。\n")
	content, _ = ReadFile(filename)
	fmt.Printf("file content: %s\n", content)
}
