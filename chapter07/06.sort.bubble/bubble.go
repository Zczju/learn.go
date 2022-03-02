package main

import (
	"fmt"
	"math/rand"
	"time"
)

func bubbleSort(arr *[]int) {
	for i := 0; i < len(*arr); i++ {
		// 计数器i,+1表示从头到尾过一遍了（走多少轮）
		for j := 0; j < len(*arr)-1-i; j++ {
			// 用来做比较的箭头j
			if (*arr)[j] > (*arr)[j+1] {
				(*arr)[j], (*arr)[j+1] = (*arr)[j+1], (*arr)[j]
			}
		}
		//fmt.Println("中间状态: ", arr)
	}
	//fmt.Println("最终状态: ", arr)
}

func main() {
	arrSize := 100000
	arr := []int{}
	for i := 0; i < arrSize; i++ {
		arr = append(arr, rand.Intn(50))
	}
	start := time.Now()
	bubbleSort(&arr)
	finish := time.Now()
	fmt.Println(finish.Sub(start))

}
