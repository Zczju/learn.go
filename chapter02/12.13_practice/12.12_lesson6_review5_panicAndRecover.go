package main

import (
	"fmt"
	"runtime/debug"
)

func panicAndRecover() {
	defer coverPanicUpgraded() // 这样能抓住严重错误，工作中最常见方式

	defer func() {
		coverPanicUpgraded() // 同样抓不住，因为已经脱离函数工作的上下文
	}()

	defer coverPanic() // 抓不住
	var nameScore map[string]int = nil

	nameScore["张宸"] = 100

}

func coverPanic() { // 未能抓住panic
	func() {
		if r := recover(); r != nil {
			fmt.Println("系统出严重故障：", r)
		}
	}() // 已经逃出了当前函数的调用过程，不在同一个栈内
}

func coverPanicUpgraded() {
	if r := recover(); r != nil {

		fmt.Println("系统出严重故障：", r)
		debug.PrintStack()
	}
}
