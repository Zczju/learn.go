package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	filePath := "C:\\Users\\Administrator\\Desktop\\ToDoList.txt"
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("无法打开文件", filePath, "错误信息是：", err)
		os.Exit(1) // 异常退出，并输出错误代码
	}
	defer file.Close()

	b := make([]byte, 10)
	var n int
	for i := 0; i < 2; i++ {
		n, err = file.Read(b) // It returns the number of bytes read and any error encountered.
		if err != nil {
			fmt.Println("无法读取文件：", err)
			os.Exit(2)
		}
		fmt.Println("n的大小：", n)
		fmt.Println("读出的内容：", string(b[:n]))
		//file.Seek(0,io.SeekStart) // seekStart表示从头开始数数，转移游标，offset表示游标移动几个byte  // 这段表示，循环读过第一次之后，将指头放回原点再读第二次；所以两次读取的是一样的
		file.Seek(3, io.SeekCurrent) // seekCurrent表示从目前游标的位置开始，转移相对位置
		//ret, err := file.Seek(0,io.SeekCurrent) // 这里的返回值ret表示当前游标的位置
		// todo handle error
	}
	//fmt.Println("读出的内容：", b) // 如果不转换string类型，输出的是一长串数字
	//fmt.Println("读出的内容：", string(b))
	//fmt.Println("n的大小：", n)
	//fmt.Println("读出的内容：", string(b[:n])) // b != b[:n] 有效数据区间（当b的长度远大于n时）
	// 一定记得给后续程序使用时，截取到实际读取到的数据，而不是全部。否则后续程序会把无效读取也作为正常数据处理
}
