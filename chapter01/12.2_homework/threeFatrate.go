package main

import "fmt"

func main() {
	var name [3]string
	var age [3]int
	var weight [3]float64
	var tall [3]float64
	var sex [3]string
	var sexWeight [3]int
	var BMI [3]float64
	var fatrate [3]float64

	fmt.Print("请输入第一个人的姓名：")
	fmt.Scanln(&name[0])
	fmt.Print("请输入第一个人的年龄（岁）：")
	fmt.Scanln(&age[0])
	fmt.Print("请输入第一个人的性别（男/女）：")
	fmt.Scanln(&sex[0])
	fmt.Print("请输入第一个人的体重（千克）：")
	fmt.Scanln(&weight[0])
	fmt.Print("请输入第一个人的身高（米）：")
	fmt.Scanln(&tall[0])

	if sex[0] == "男" {
		sexWeight[0] = 1
	} else {
		sexWeight[0] = 0
	}
	BMI[0] = weight[0] / (tall[0] * tall[0])
	fatrate[0] = (1.2*BMI[0] + 0.23*float64(age[0]) - 5.4 - 10.8*float64(sexWeight[0])) / 100

	fmt.Println("您的")

}
