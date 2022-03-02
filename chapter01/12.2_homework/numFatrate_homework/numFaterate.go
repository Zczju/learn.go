package main

import "fmt"

func main() {
	const NumsPeople = 3
	// 定义输入人数

	var name [NumsPeople]string
	var age [NumsPeople]int
	var weight [NumsPeople]float64
	var tall [NumsPeople]float64
	var sex [NumsPeople]string
	var sexWeight [NumsPeople]int
	var BMI [NumsPeople]float64
	var fatrate [NumsPeople]float64
	var sum float64
	//定义输入输出变量

	i := 0
	for {

		fmt.Print("请输入姓名：")
		fmt.Scanln(&name[i])
		fmt.Print("请输入年龄（岁）：")
		fmt.Scanln(&age[i])
		fmt.Print("请输入性别（男/女）：")
		fmt.Scanln(&sex[i])
		fmt.Print("请输入体重（千克）：")
		fmt.Scanln(&weight[i])
		fmt.Print("请输入身高（米）：")
		fmt.Scanln(&tall[i])
		//输入

		if sex[i] == "男" {
			sexWeight[i] = 1
		} else {
			sexWeight[i] = 0
		}
		// 确定参数值

		BMI[i] = weight[i] / (tall[i] * tall[i])
		fatrate[i] = (1.2*BMI[i] + 0.23*float64(age[i]) - 5.4 - 10.8*float64(sexWeight[i])) / 100
		// 计算BMI和体脂率

		fmt.Println("您的BMI是：", BMI[i])
		fmt.Println("体脂率是：", fatrate[i])
		//输出BMI和体脂率

		sum = sum + fatrate[i]
		// 对体脂率求和

		if sex[i] == "男" {
			if age[i] >= 18 && age[i] <= 39 {
				if fatrate[i] <= 0.1 {
					fmt.Println("目前状况是：偏瘦；要多吃肉哦！")
				} else if fatrate[i] > 0.1 && fatrate[i] <= 0.16 {
					fmt.Println("目前状况是：标准；请继续保持！")
				} else if fatrate[i] > 0.16 && fatrate[i] <= 0.21 {
					fmt.Println("目前状况是：偏重；少吃点长点心吧！")
				} else if fatrate[i] > 0.21 && fatrate[i] <= 0.26 {
					fmt.Println("目前状况是：肥胖；不能再吃了哥！")
				} else {
					fmt.Println("目前状况是：严重肥胖；该减减肥了！")
				}

			} else if age[i] >= 40 && age[i] <= 59 {
				if fatrate[i] <= 0.11 {
					fmt.Println("目前状况是：偏瘦；要多吃肉哦！")
				} else if fatrate[i] >= 0.11 && fatrate[i] <= 0.17 {
					fmt.Println("目前状况是：标准；请继续保持！")
				} else if fatrate[i] >= 0.17 && fatrate[i] <= 0.22 {
					fmt.Println("目前状况是：偏重；少吃点长点心吧！")
				} else if fatrate[i] >= 0.22 && fatrate[i] <= 0.27 {
					fmt.Println("目前状况是：肥胖；不能再吃了哥！")
				} else {
					fmt.Println("目前状况是：严重肥胖；该减减肥了！")
				}

			} else if age[i] >= 60 {
				if fatrate[i] <= 0.13 {
					fmt.Println("目前状况是：偏瘦；要多吃肉哦！")
				} else if fatrate[i] >= 0.13 && fatrate[i] <= 0.19 {
					fmt.Println("目前状况是：标准；请继续保持！")
				} else if fatrate[i] >= 0.19 && fatrate[i] <= 0.24 {
					fmt.Println("目前状况是：偏重；少吃点长点心吧！")
				} else if fatrate[i] >= 0.24 && fatrate[i] <= 0.29 {
					fmt.Println("目前状况是：肥胖；不能再吃了哥！")
				} else {
					fmt.Println("目前状况是：严重肥胖；该减减肥了！")
				}
			} else {
				fmt.Println("小孩子没有体脂，请长大点再来吧")
			}
		}
		//给出男性的健康建议

		if sex[i] == "女" {
			if age[i] >= 18 && age[i] <= 39 {
				if fatrate[i] <= 0.2 {
					fmt.Println("目前状况是：偏瘦；要多吃肉哦！")
				} else if fatrate[i] > 0.2 && fatrate[i] <= 0.27 {
					fmt.Println("目前状况是：标准；请继续保持！")
				} else if fatrate[i] > 0.27 && fatrate[i] <= 0.34 {
					fmt.Println("目前状况是：偏重；少吃点长点心吧！")
				} else if fatrate[i] > 0.34 && fatrate[i] <= 0.39 {
					fmt.Println("目前状况是：肥胖；不能再吃了姐！")
				} else {
					fmt.Println("目前状况是：严重肥胖；该减减肥了！")
				}

			} else if age[i] >= 40 && age[i] <= 59 {
				if fatrate[i] <= 0.21 {
					fmt.Println("目前状况是：偏瘦；要多吃肉哦！")
				} else if fatrate[i] >= 0.21 && fatrate[i] <= 0.28 {
					fmt.Println("目前状况是：标准；请继续保持！")
				} else if fatrate[i] >= 0.28 && fatrate[i] <= 0.35 {
					fmt.Println("目前状况是：偏重；少吃点长点心吧！")
				} else if fatrate[i] >= 0.35 && fatrate[i] <= 0.40 {
					fmt.Println("目前状况是：肥胖；不能再吃了姐！")
				} else {
					fmt.Println("目前状况是：严重肥胖；该减减肥了！")
				}

			} else if age[i] >= 60 {
				if fatrate[i] <= 0.22 {
					fmt.Println("目前状况是：偏瘦；要多吃肉哦！")
				} else if fatrate[i] >= 0.22 && fatrate[i] <= 0.29 {
					fmt.Println("目前状况是：标准；请继续保持！")
				} else if fatrate[i] >= 0.29 && fatrate[i] <= 0.36 {
					fmt.Println("目前状况是：偏重；少吃点长点心吧！")
				} else if fatrate[i] >= 0.36 && fatrate[i] <= 0.41 {
					fmt.Println("目前状况是：肥胖；不能再吃了姐！")
				} else {
					fmt.Println("目前状况是：严重肥胖；该减减肥了！")
				}
			} else {
				fmt.Println("小孩子没有体脂，请长大点再来吧")
			}
		}
		// 给出女性的健康建议

		i++
		if i > NumsPeople-1 {
			fmt.Println("输入完毕")
			break
		}
		//..

		var whetherContinue string
		fmt.Print("是否录入下一个？(y/n)")
		fmt.Scanln(&whetherContinue)
		if whetherContinue != "y" {
			break
		}
		// 录入提醒
	}
	fmt.Println("总人数为：", NumsPeople)
	fmt.Println("平均体脂率为：", sum/NumsPeople)
}
