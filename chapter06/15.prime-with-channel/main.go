package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	// 计算素数，实现负载均衡，平均分配资源
	workerNumber := 8
	maxNum := 200000
	baseNumCh := make(chan int, 1024)
	result := make(chan int, maxNum)
	wg := sync.WaitGroup{}
	start := time.Now()

	wg.Add(workerNumber)
	for i := 0; i < workerNumber; i++ {
		go func() {
			defer wg.Done()
			for oNum := range baseNumCh {
				if isPrime(oNum) {
					result <- oNum
				}
			}
		}()
	}
	for num := 2; num <= maxNum; num++ {
		baseNumCh <- num
	}
	close(baseNumCh)
	wg.Wait()

	finish := time.Now()
	fmt.Println(len(result))
	fmt.Println("共耗时：", finish.Sub(start))

}

func isPrime(num int) (isPrime bool) {
	for i := 2; i < num; i++ {
		if num%i == 0 {
			isPrime = false
			return
		}
	}
	isPrime = true
	return
}
