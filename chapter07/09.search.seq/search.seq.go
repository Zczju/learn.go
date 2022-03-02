package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"time"
)

var totalCompare int

func main() {
	//size := 1000
	arr := sampleData
	//generateRandomData(size)

	// 1. 501在不在里边？
	// 2. 888在不在？
	// 3. 900?
	// 4. 3?
	// ...
	startTime := time.Now()
	for i := 0; i < 1000000; i++ { //总用时:  6.9954001s 总比较次数:  3767000000
		search(&arr, 501)
		search(&arr, 888)
		search(&arr, 900)
		search(&arr, 3)
	}

	finishTime := time.Now()
	fmt.Println("总用时: ", finishTime.Sub(startTime))
	fmt.Println("总比较次数: ", totalCompare)

}

func generateRandomData(size int) []int64 {
	arr := make([]int64, 0, size)

	for i := 0; i < size; i++ {
		// 使用crypto/rand包实现了真随机数
		num, _ := rand.Int(rand.Reader, big.NewInt(3000)) // num是big.int类型,不能直接append
		arr = append(arr, num.Int64())
	}
	return arr
}
func search(arrP *[]int64, targetNum int64) bool {
	for _, num := range *arrP {
		totalCompare++
		if num == targetNum {
			return true
		}
	}
	return false
}
