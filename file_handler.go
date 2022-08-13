/**
 * @Time : 2022/8/11 6:53 PM
 * @Author : frankj
 * @Email : --
 * @Description : --
 * @Revise : --
 */

package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

// read file
func read(path string) (content string) {
	filePath := path
	file, err := os.OpenFile(filePath, os.O_RDWR|os.O_APPEND, 0666)
	if err != nil {
		fmt.Printf("文件打开失败: %v", err)
		return
	}
	//及时关闭file句柄
	defer file.Close()
	//读原来文件的内容，并且显示在终端
	reader := bufio.NewReader(file)
	for {
		str, err := reader.ReadString('\n')
		content += str
		if err == io.EOF {
			break
		}
	}
	return
}

// write new file
func write(path, content string) {
	file, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		fmt.Printf("open file failed: %v \n", err)
		return
	}
	defer file.Close()
	writer := bufio.NewWriter(file)

	_, err = writer.WriteString(content)
	if err != nil {
		fmt.Printf("writer write string failed: %v \n", err)
		return
	}

	err = writer.Flush()
	if err != nil {
		fmt.Printf("writer flush failed: %v \n", err)
	}
}

// write append file
func writeAppend(path, content string) {
	file, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Printf("open file failed: %v \n", err)
		return
	}
	defer file.Close()
	write := bufio.NewWriter(file)
	_, err = write.WriteString(content)
	if err != nil {
		fmt.Printf("writer write string append failed: %v \n", err)
		return
	}

	err = write.Flush()
	if err != nil {
		fmt.Printf("writer flush failed: %v \n", err)
	}
}

// determine whether the file exists
func isExist(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		if os.IsNotExist(err) {
			return false
		}
		return false
	}
	return true
}
