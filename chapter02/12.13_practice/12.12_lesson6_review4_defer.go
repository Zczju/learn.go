package main

import (
	"fmt"
	"os"
	"time"
)

func openFile() {
	filename := "/"
	fileObj, err := os.Open(filename)
	if err != nil {
		fmt.Println("error:", err)
		os.Exit(1)
	}
	defer fileObj.Close() //defer在函数运行结束后执行，若与前面共享变量，则与共享变量产生关系

	defer func() {
		fileObj.Close()
	}()
	fileObj = nil

}

func deferGuess() {
	startTime := time.Now() //获取当前时间

	defer fmt.Println("duration:", time.Now().Sub(startTime)) // 计算持续时间
	// defer运行只运行最后一层；Println是最后运行的，里面的东西会预先运行好

	defer func() {
		fmt.Println("duration upgraded:", time.Now().Sub(startTime))
	}()
	// 直接准备好的是函数，而函数在最后运行

	time.Sleep(5 * time.Second)
	fmt.Println("start time:", startTime)
	fmt.Println("finish time:", time.Now())

}
