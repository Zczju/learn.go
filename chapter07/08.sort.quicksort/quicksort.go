package main

import (
	"fmt"
	"math/rand"
	"time"
)

func quickSort(arr *[]int, start, end int) {
	pivotIdx := (start + end) / 2
	pivotVal := (*arr)[pivotIdx]
	l, r := start, end
	for l <= r {
		for (*arr)[l] < pivotVal {
			l++
		}
		for (*arr)[r] > pivotVal {
			r--
		}
		if l >= r {
			break
		}
		(*arr)[l], (*arr)[r] = (*arr)[r], (*arr)[l]
		l++
		r--
	}
	if l == r {
		l++
		r--
	}
	if r > start {
		quickSort(arr, start, r)
	}
	if l < end {
		quickSort(arr, l, end)
	}

}

func main() {
	arrSize := 100000
	arr := []int{}
	for i := 0; i < arrSize; i++ {
		arr = append(arr, rand.Intn(50))
	}

	start := time.Now()
	quickSort(&arr, 0, arrSize-1)
	finish := time.Now()
	fmt.Println(finish.Sub(start))

}
