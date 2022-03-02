package main

import (
	"fmt"
)

func main() {
	/*
		{
			// 计算器实现案例一
			var input string
			fmt.Scanln(&input)
			pieces := strings.Split(input, "&") // strings.Split 输入s和sep，以sep分割s字符串，并装入一个切片内返回
			switch pieces[1] {
			case "+":
				left, _ := strconv.Atoi(pieces[0]) // 字符串转整型，两个返回值，第二个err用来判断转换是否成功
				right, _ := strconv.Atoi(pieces[2])
				fmt.Println(left + right)
			case "-":
				left, _ := strconv.Atoi(pieces[0])
				right, _ := strconv.Atoi(pieces[2])
				fmt.Println(left - right)
			case "*":
				left, _ := strconv.Atoi(pieces[0])
				right, _ := strconv.Atoi(pieces[2])
				fmt.Println(left * right)
			case "/":
				left, _ := strconv.Atoi(pieces[0])
				right, _ := strconv.Atoi(pieces[2])
				fmt.Println(left / right)
			case "%":
				left, _ := strconv.Atoi(pieces[0])
				right, _ := strconv.Atoi(pieces[2])
				fmt.Println(left % right)
			default:
				fmt.Println("not supported calculate rule")
			}
		}
	*/

	// 计算器实现案例2
	var left, right int = 1, 2
	// var op string = "+"
	/*
		fmt.Scanln(&left)
		fmt.Print(" ")
		fmt.Scanln(&op)
		fmt.Print(" ")
		fmt.Scanln(&right)
		switch op {
		case "+":
			fmt.Println(left + right)
		case "-":
			fmt.Println(left - right)14
		case "*":
			fmt.Println(left * right)
		case "/":
			fmt.Println(left / right)
		case "%":
			fmt.Println(left % right)
		default:
			fmt.Println("not supported calculate rule")
		}
	*/

	c := Calculator{
		left:  left,
		right: right,
		// op: op,  // 直接调了具体的函数，就不需要op了
	}
	fmt.Printf("&c @ main = %p\n", &c)
	fmt.Println(c.Add())
	fmt.Println("c.result = ", c.result)

	// newC := NewCalculator{}
	newC := NewCalculator{&Calculator{}}
	// 若继承的是指针，必须将指针实例化之后才可使用
	newC.left = 100
	newC.right = 200
	fmt.Println(newC.Add())

	mc := MyCommand{commandOptions: map[string]string{}}
	mc.commandOptions["aa"] = "AAA"
	fmt.Println(mc.ToCmdStr())

}

type MyCommand struct {
	mainCommand    *string // 必须指向一个变量后才能进行后续的运算
	commandOptions map[string]string
}

// 定义对象的时候，对象的每一个成员在用之前也是需要实例化的
// 不论嵌套还是成员，都需要实例化

func (my MyCommand) ToCmdStr() string {
	out := ""
	for k, v := range my.commandOptions {
		out = out + fmt.Sprintf("--%s=%s", k, v)
	}
	return out
}
