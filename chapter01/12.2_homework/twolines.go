package main

import (
	"fmt"
)

func main() {
	for {
		fmt.Println("现在是第一条直线：")
		var x1 float64
		fmt.Print("请输入第一个点的横坐标：")
		fmt.Scanln(&x1)
		var y1 float64
		fmt.Print("请输入第一个点的纵坐标：")
		fmt.Scanln(&y1)

		var x2 float64
		fmt.Print("请输入第二个点的横坐标：")
		fmt.Scanln(&x2)
		var y2 float64
		fmt.Print("请输入第二个点的纵坐标：")
		fmt.Scanln(&y2)

		var k1 float64 = (y2 - y1) / (x2 - x1)

		fmt.Println("现在是第二条直线")
		var x3 float64
		fmt.Print("请输入第一个点的横坐标：")
		fmt.Scanln(&x3)
		var y3 float64
		fmt.Print("请输入第一个点的纵坐标：")
		fmt.Scanln(&y3)

		var x4 float64
		fmt.Print("请输入第二个点的横坐标：")
		fmt.Scanln(&x4)
		var y4 float64
		fmt.Print("请输入第二个点的纵坐标：")
		fmt.Scanln(&y4)

		var k2 float64 = (y4 - y3) / (x4 - x3)
		if k1 == k2 {
			fmt.Println("两条直线平行")
		} else {
			fmt.Println("两条直线不平行")
		}

		var whetherContinue string
		fmt.Println("是否继续输入?(y/n)")
		fmt.Scanln(&whetherContinue)
		if whetherContinue != "y" {
			break
		}
	}
}
