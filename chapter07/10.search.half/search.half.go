package main

import (
	"fmt"
	"time"
)

var totalCompare int

func main() {
	//size := 1000
	//arr := sampleData
	//generateRandomData(size)

	// 1. 501在不在里边？
	// 2. 888在不在？
	// 3. 900?
	// 4. 3?
	// ...

	// 复制一份新数组，避免修改原始数组
	newArr := append([]int64{}, sampleData...)
	startTime := time.Now()
	// 排序
	quickSort(&newArr, 0, len(newArr)-1)

	for i := 0; i < 100000; i++ { // 总用时:  12.0007ms 总比较次数:  4200000
		search(&newArr, 501)
		search(&newArr, 888)
		search(&newArr, 900)
		search(&newArr, 3)
	}

	finishTime := time.Now()
	fmt.Println("总用时: ", finishTime.Sub(startTime))
	fmt.Println("总比较次数: ", totalCompare)

}

func search(arrP *[]int64, targetNum int64) bool {
	return searchHalf(arrP, 0, len(*arrP)-1, targetNum)
}

func searchHalf(arrP *[]int64, left, right int, targetNum int64) bool {
	middleIndex := (left + right) / 2
	halfData := (*arrP)[middleIndex]

	totalCompare++ // 增加计数器数值

	if targetNum > halfData {
		if left == right {
			return false
		}
		return searchHalf(arrP, middleIndex+1, right, targetNum)
	} else if targetNum < halfData {
		if left == right {
			return false
		}
		return searchHalf(arrP, left, middleIndex, targetNum) // todo 数组越界问题 !
	} else {
		return true
	}
}
