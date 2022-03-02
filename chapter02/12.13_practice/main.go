package main

import (
	"fmt"
	"time"
)

var tall float64
var weight float64

// 全局变量，在外面和里面都有效

func main() {
	//guessNum(1, 100)

	panicAndRecover()

	deferGuess()
	fmt.Println(" sleep somewhile")
	time.Sleep(10 * time.Second)

	openFile()
	fmt.Println(" sleep somewhile")
	time.Sleep(10 * time.Second)

	closureMain()
	fmt.Println(" sleep somewhile")
	time.Sleep(10 * time.Second)

	// close call
	fmt.Println(calcSum(13, 3, 43, 56, 64, 67, 67, 2, 5, 57))
	fmt.Println(calcSum(13, 3, 43, 56, 64, 67, 67, 2, 5, 57))
	fmt.Println(calcSum(13, 3, 43, 56, 64, 67, 67, 2, 5, 57))
	fmt.Println(calcSum(13, 3, 43, 56, 64, 67, 67, 2, 5, 57))
	fmt.Println(calcSum(13, 3, 43, 56, 64, 67, 67, 2, 5, 57))
	fmt.Println(calcSum(13, 3, 43, 56, 64, 67, 67, 2, 5, 57))
	fmt.Println(calcSum(13, 3, 43, 56, 64, 67, 67, 2, 5, 57))

	showUsedTimes()

	// 递归：函数调用链中存在自己调用自己
	// 需要一个合适的终止条件
	// 例子 ：汉诺塔，斐波那契
	fmt.Println(fib(1)) // 100会算很久 会算2的100次方次
	// 解决方式 ： 递归换成循环

	fmt.Println("done calc, sleep somewhile")
	time.Sleep(10 * time.Second)

	//sampleSubdomain()
	fmt.Println("全局变量赋值前：")
	calcAdd() // ??  // 0
	tall, weight = 1.70, 70.0
	fmt.Println("全局变量赋值后：")
	calcAdd() // ?? // 71.7

	tall, weight := 100.0, 70.0 // 局部变量仅在main函数作用域内有效
	fmt.Println(tall, weight)
	calcAdd()                   // ??   //71.7
	tall, weight = 200.00, 70.0 // 就近原则，赋值赋给局部变量
	calcAdd()                   //?? //71.7

	calculatorAdd := func(a, b float64) float64 {
		return a + b
	}
	// fmt.Println(a,b) a,b只在花括号内有效
	result := calculatorAdd(1, 3)
	fmt.Println(result)

	{
		//fmt.Scanln....
		personTall := 1.81
		personWeight := 90.0
		calculatorAdd(personTall, personWeight)
		//suggestions
	}

	{
		//fmt.Scanln....
		personTall := 1.81
		personWeight := 90.0
		calculatorAdd(personTall, personWeight)
		//suggestions
	}
	// 同名变量在同一个函数内存在，可分别放入不同的作用域内，用{}隔开
	// 作用范围只在两个花括号内部，外部无效
	// fmt.Println(personTall)

}

func guessNum(left, right uint) {
	guessNumber := (left + right) / 2
	var getFromInput int
	fmt.Println("我猜是 ：", guessNumber)
	fmt.Print("如果高了，输入1；如果低了，输入0；如果猜对了，输入9 ：")
	fmt.Scanln(&getFromInput)
	switch getFromInput {
	case 1:
		{
			if left == right {
				fmt.Println("你是不是改变主意了？")
				return
			}
			guessNum(left, guessNumber-1)
		}
	case 0:
		{
			if left == right {
				fmt.Println("你是不是改变主意了？")
				return
			}
			guessNum(guessNumber+1, right)
		}
	case 9:
		{
			fmt.Println("你心里想的数字是：", guessNumber)
		}

	}
}
