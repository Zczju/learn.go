package main

import "fmt"

func main() {
	{
		a := "hello"
		fmt.Println(a)

		aBytes := []byte(a)
		fmt.Println(aBytes)

		fmt.Println("修改切片内的内容")
		aBytes[0] = 'H'
		a = string(aBytes) // 将切片转回string类型
		fmt.Println(a)
	}

	{
		a := "您好"
		aBytes := []byte(a)
		fmt.Println(aBytes)
		aBytes[0] = 'H'
		a = string(aBytes)
		fmt.Println(a)
	}
	{
		a := "您好"
		aBytes := []rune(a)
		fmt.Println(aBytes)
		aBytes[0] = 'H'
		a = string(aBytes)
		fmt.Println(a)
	}

}
