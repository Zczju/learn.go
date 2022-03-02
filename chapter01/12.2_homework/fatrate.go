package main

import (
	"fmt"
)

func main() {
	for {
		//录入信息
		weight, tall, age, sexWeight, sex := getInfosInput()

		//计算体脂率
		fatrate := calcFatrate(weight, tall, age, sexWeight)

		//给出健康建议
		giveHealthySuggestions(sex, age, fatrate)

		var whetherContinue string
		fmt.Print("是否录入下一个？(y/n)")
		fmt.Scanln(&whetherContinue)
		if whetherContinue != "y" {
			break
		}
	}
}

func giveHealthySuggestions(sex string, age int, fatrate float64) {
	if sex == "女" {
		//编写女性的体脂率与体脂状态表
		if age >= 18 && age <= 39 {
			if fatrate <= 0.2 {
				fmt.Println("目前状况是：偏瘦；要多吃肉哦！")
			} else if fatrate > 0.2 && fatrate <= 0.27 {
				fmt.Println("目前状况是：标准；请继续保持！")
			} else if fatrate > 0.27 && fatrate <= 0.34 {
				fmt.Println("目前状况是：偏重；少吃点长点心吧！")
			} else if fatrate > 0.34 && fatrate <= 0.39 {
				fmt.Println("目前状况是：肥胖；不能再吃了姐！")
			} else {
				fmt.Println("目前状况是：严重肥胖；该减减肥了！")
			}

		} else if age >= 40 && age <= 59 {
			if fatrate <= 0.21 {
				fmt.Println("目前状况是：偏瘦；要多吃肉哦！")
			} else if fatrate >= 0.21 && fatrate <= 0.28 {
				fmt.Println("目前状况是：标准；请继续保持！")
			} else if fatrate >= 0.28 && fatrate <= 0.35 {
				fmt.Println("目前状况是：偏重；少吃点长点心吧！")
			} else if fatrate >= 0.35 && fatrate <= 0.40 {
				fmt.Println("目前状况是：肥胖；不能再吃了姐！")
			} else {
				fmt.Println("目前状况是：严重肥胖；该减减肥了！")
			}

		} else if age >= 60 {
			if fatrate <= 0.22 {
				fmt.Println("目前状况是：偏瘦；要多吃肉哦！")
			} else if fatrate >= 0.22 && fatrate <= 0.29 {
				fmt.Println("目前状况是：标准；请继续保持！")
			} else if fatrate >= 0.29 && fatrate <= 0.36 {
				fmt.Println("目前状况是：偏重；少吃点长点心吧！")
			} else if fatrate >= 0.36 && fatrate <= 0.41 {
				fmt.Println("目前状况是：肥胖；不能再吃了姐！")
			} else {
				fmt.Println("目前状况是：严重肥胖；该减减肥了！")
			}
		} else {
			fmt.Println("小孩子没有体脂，请长大点再来吧")
		}

	}
	if sex == "男" {
		if age >= 18 && age <= 39 {
			if fatrate <= 0.1 {
				fmt.Println("目前状况是：偏瘦；要多吃肉哦！")
			} else if fatrate > 0.1 && fatrate <= 0.16 {
				fmt.Println("目前状况是：标准；请继续保持！")
			} else if fatrate > 0.16 && fatrate <= 0.21 {
				fmt.Println("目前状况是：偏重；少吃点长点心吧！")
			} else if fatrate > 0.21 && fatrate <= 0.26 {
				fmt.Println("目前状况是：肥胖；不能再吃了哥！")
			} else {
				fmt.Println("目前状况是：严重肥胖；该减减肥了！")
			}

		} else if age >= 40 && age <= 59 {
			if fatrate <= 0.11 {
				fmt.Println("目前状况是：偏瘦；要多吃肉哦！")
			} else if fatrate >= 0.11 && fatrate <= 0.17 {
				fmt.Println("目前状况是：标准；请继续保持！")
			} else if fatrate >= 0.17 && fatrate <= 0.22 {
				fmt.Println("目前状况是：偏重；少吃点长点心吧！")
			} else if fatrate >= 0.22 && fatrate <= 0.27 {
				fmt.Println("目前状况是：肥胖；不能再吃了哥！")
			} else {
				fmt.Println("目前状况是：严重肥胖；该减减肥了！")
			}

		} else if age >= 60 {
			if fatrate <= 0.13 {
				fmt.Println("目前状况是：偏瘦；要多吃肉哦！")
			} else if fatrate >= 0.13 && fatrate <= 0.19 {
				fmt.Println("目前状况是：标准；请继续保持！")
			} else if fatrate >= 0.19 && fatrate <= 0.24 {
				fmt.Println("目前状况是：偏重；少吃点长点心吧！")
			} else if fatrate >= 0.24 && fatrate <= 0.29 {
				fmt.Println("目前状况是：肥胖；不能再吃了哥！")
			} else {
				fmt.Println("目前状况是：严重肥胖；该减减肥了！")
			}
		} else {
			fmt.Println("小孩子没有体脂，请长大点再来吧")
		}
	}
}

func getInfosInput() (float64, float64, int, int, string) {
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
	fmt.Print("年龄（岁）：")
	fmt.Scanln(&age)

	var sexWeight int
	var sex string
	fmt.Print("性别（男/女）：")
	fmt.Scanln(&sex)

	if sex == "男" {
		sexWeight = 1
	} else {
		sexWeight = 0
	}
	return weight, tall, age, sexWeight, sex
}

func calcFatrate(weight float64, tall float64, age int, sexWeight int) float64 {
	// 计算体脂率
	bmi := weight / (tall * tall)
	fmt.Println("您的BMI是：", bmi)
	fatRate := (1.2*bmi + 0.23*float64(age) - 5.4 - 10.8*float64(sexWeight)) / 100
	fmt.Println("体脂率是：", fatRate)
	return fatRate
}
