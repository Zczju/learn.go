package main

import "fmt"

func calcAdd() float64 {
	fmt.Println(tall + weight)
	// 局部变量，仅在{}内有效
	return 0
}

func sampleSubdomain() {
	name := "小强" // 声明变量name，值是"小强"
	fmt.Println("名字是：", name)
	{
		name = "小强-Upgraded" //  作用域内可对外部的变量(全局变量)赋值
		fmt.Println("名字是：", name)
		name = "kr" // 注意这里是=等号，是对括号外变量的重新赋值，把值给了name，值是"kr"
		// name := "kr" //冒号引号才是重新定义
		fmt.Println("名字是：", name)
	}
	fmt.Println(">>>>>>名字是：", name) // 小强？？ kr？？

	if name == "小强" {
		a := 3
		fmt.Println(a)
	} else {
		a := 4
		fmt.Println(a)
	} // a为if作用域内变量，仅在作用域内有效
}

func sampleSubdomainif() {
	if v := calcAdd(); v == 0 {
		fmt.Println(v)
	} else {
		fmt.Println(v)
	}
	// fmt.Println(v)  // 无效。 v的有效范围为if-else block

}

func sampleSubdomainfor() {
	for i := 0; i < 10; i++ {
		fmt.Println("hello golang", i)
	}
	// fmt.Println(i)  // 无效. i的有效范围为for block
}
