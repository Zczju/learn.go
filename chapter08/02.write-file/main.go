package main

import (
	"fmt"
	"os"
)

func main() {
	filePath := "C:\\Users\\Administrator\\Desktop\\ToDoList-new.txt"
	//writeFile(filePath)
	//file.Sync() // 告诉操作系统直接把内容都写到磁盘上
	writeFileWithAppend(filePath)
	// os: invalid use of WriteAt on file opened with O_APPEND
}

func writeFile(filePath string) {
	file, err := os.Create(filePath)
	if err != nil {
		fmt.Println("无法打开文件", filePath, "错误信息是：", err)
		os.Exit(1) // 异常退出，并输出错误代码
	}
	defer file.Close()

	_, err = file.Write([]byte("this is first line---"))
	fmt.Println(err)
	_, err = file.Write([]byte("this is first line\n"))
	// todo handle error
	fmt.Println(err)
	_, err = file.WriteString("第二行内容\n")
	// todo handle error
	fmt.Println(err)
	_, err = file.WriteAt([]byte("CHANGED"), 0)
	// todo handle error
	fmt.Println(err)
}

func writeFileWithAppend(filePath string) {
	file, err := os.OpenFile(filePath, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0777) // linux file permission // 自己，同组人，所有人
	if err != nil {
		fmt.Println("无法打开文件", filePath, "错误信息是：", err)
		os.Exit(1) // 异常退出，并输出错误代码
	}
	defer file.Close()

	_, err = file.Write([]byte("this is first line---"))
	fmt.Println(err)
	_, err = file.Write([]byte("this is first line\n"))
	// todo handle error
	fmt.Println(err)
	_, err = file.WriteString("第二行内容\n")
	// todo handle error
	fmt.Println(err)
	_, err = file.WriteAt([]byte("CHANGED"), 0) // O_APPEND方式只能在尾部追加，不能修改
	// todo handle error
	fmt.Println(err)
}
