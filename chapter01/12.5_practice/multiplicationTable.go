package main

import "fmt"

func main() {
	var num []int
	for i := 1; i < 10; i++ {
		num = append(num, i)
	}

	for j := 1; j < 10; j++ {
		fmt.Printf("%d*%d\n", j, num[j-1])
	}

}
