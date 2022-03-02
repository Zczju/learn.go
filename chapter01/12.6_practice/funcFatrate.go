package main

import "fmt"

const num = 3

var name [num]string
var age [num]int
var sex [num]string
var weight [num]float64
var tall [num]float64
var BMI [num]float64
var fatrate [num]float64
var sexWeight [num]int

func main() {
	i := 0
	for {
		getInfosInput(name[i], age[i], sex[i], tall[i], weight[i])

		if sex[i] == "男" {
			sexWeight[i] = 1
		} else {
			sexWeight[i] = 0
		}

		BMI[i] = calcBMI(tall[i], weight[i])
		fmt.Println("您的BMI是：", BMI[i])

		fatrate[i] = calcFatRate(BMI[i], age[i], sexWeight[i])
		fmt.Println("您的体脂率是：", fatrate[i])

	}

}

func getInfosInput(name string, age int, sex string, tall float64, weight float64) {
	fmt.Print("姓名：")
	fmt.Scanln(&name)

	fmt.Print("年龄（岁）：")
	fmt.Scanln(&age)

	fmt.Print("性别（男/女）：")
	fmt.Scanln(&sex)

	fmt.Print("身高（米）：")
	fmt.Scanln(&tall)

	fmt.Print("体重（千克）：")
	fmt.Scanln(&weight)

}
func calcBMI(tall float64, weight float64) float64 {
	return weight / (tall * tall)
}

func calcFatRate(BMI float64, age int, sexWeight int) float64 {
	return (1.2*BMI + 0.23*float64(age) - 5.4 - 10.8*float64(sexWeight)) / 100
}
