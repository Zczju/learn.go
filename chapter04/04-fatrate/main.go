package main

import "fmt"

func main() {
	/*
		实现案例1: 人与计算器分离的体脂服务(两个对象)
		person := getFakePersonFromInput() // 获取个人信息
		c:= Calc{}
		c.BMI(person)
		c.FatRate(person) // 计算
		fmt.Println(person)
		sugg := fatRateSuggestions{}
		suggestion := sugg.GetSuggestions(person) // 拿建议
		fmt.Println(suggestion)
	*/

	// 案例二 人完成计算的体脂服务
	frSvc := &fatRateService{s: GetFatRateSuggestion()}

	fakePerson := getFakePersonFromInput()
	fmt.Println(frSvc.GiveSuggestionToPerson(fakePerson))

	for {
		p := getPersonInfoFromInput()
		fmt.Println(frSvc.GiveSuggestionToPerson(p))
	}
}

func getPersonInfoFromInput() *Person {
	// 录入各项
	var name string
	fmt.Print("姓名：")
	fmt.Scanln(&name)

	var weight float64
	fmt.Print("体重（千克）：")
	fmt.Scanln(&weight)

	var tall float64
	fmt.Print("身高（米）：")
	fmt.Scanln(&tall)
	var age int
	fmt.Print("年龄：")
	fmt.Scanln(&age)

	sex := "男"
	fmt.Print("性别（男/女）：")
	fmt.Scanln(&sex)

	return &Person{
		name:   name,
		age:    age,
		weight: weight,
		tall:   tall,
		sex:    sex,
	}
}

func getFakePersonFromInput() *Person {
	return &Person{
		name:   "小强",
		sex:    "男",
		tall:   1.7,
		weight: 70,
		age:    35,
	}
}
