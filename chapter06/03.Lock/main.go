package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	for i := 0; i < 1; i++ {
		// countDict()
		// countDictGoRoutineSafe()
		// countDictForgetUnlock()
		// countDictLockPrice()
		countDictGoRoutineSafeRW()
	}

}

func countDict() {
	fmt.Println("开始数")
	var totalCount int64 = 0
	wg := sync.WaitGroup{}
	wg.Add(5000)
	for p := 0; p < 5000; p++ {
		go func() {
			defer wg.Done()
			// fmt.Print("正在统计第", p, "页，")
			totalCount += 100 // totalCount = totalCount + 100 // 注意：这里有重复覆盖的问题
		}()
	}
	wg.Wait()
	fmt.Println("预计有", 100*5000, "字")
	fmt.Println("总共有", totalCount, "字")
}

func countDictGoRoutineSafe() {
	fmt.Println("开始数")
	var totalCount int64 = 0
	totalCountLock := sync.Mutex{} // 变量名+Lock 清晰的知道锁了什么

	wg := sync.WaitGroup{}
	wg.Add(5000)
	for p := 0; p < 5000; p++ {
		go func() {
			defer wg.Done()
			totalCountLock.Lock()
			// fmt.Print("正在统计第", p, "页，")
			defer totalCountLock.Unlock()
			totalCount += 100 // totalCount = totalCount + 100 // 注意：这里有重复覆盖的问题
			// totalCountLock.Unlock() 也可以用defer来解锁
		}()
	}
	wg.Wait()
	fmt.Println("预计有", 100*5000, "字")
	fmt.Println("总共有", totalCount, "字")
}
func countDictGoRoutineSafeRW() {
	fmt.Println("开始数")
	var totalCount int64 = 0
	totalCountLock := sync.RWMutex{} // 改用读写锁

	wg := sync.WaitGroup{}
	workerCount := 5
	wg.Add(workerCount)

	doneCh := make(chan struct{})
	for p := 0; p < workerCount; p++ {
		go func(p int) { // 读锁可以多个go routine 同时拿到
			fmt.Println(p, "读锁开始时间:", time.Now())
			totalCountLock.RLock()
			fmt.Println(p, "读锁拿到锁时间:", time.Now())
			time.Sleep(1 * time.Second)
			totalCountLock.RUnlock()
			fmt.Println(p, "读锁释放锁时间:", time.Now())
		}(p)
	}
	for p := 0; p < workerCount; p++ {
		go func(p int) {
			defer wg.Done()

			fmt.Println(p, "写锁开始时间:", time.Now())
			totalCountLock.Lock()
			fmt.Println(p, "写锁拿到锁时间:", time.Now())
			// fmt.Print("正在统计第", p, "页，")
			defer totalCountLock.Unlock()
			totalCount += 100 // totalCount = totalCount + 100 // 注意：这里有重复覆盖的问题
			// totalCountLock.Unlock() 也可以用defer来解锁
		}(p)
	}
	wg.Wait()
	close(doneCh)
	time.Sleep(1 * time.Second)
	fmt.Println("预计有", 100*5000, "字")
	fmt.Println("总共有", totalCount, "字")
}

func countDictForgetUnlock() {
	fmt.Println("开始数")
	var totalCount int64 = 0
	totalCountLock := sync.Mutex{} // 变量名+Lock 清晰的知道锁了什么

	wg := sync.WaitGroup{}
	wg.Add(5)
	for p := 0; p < 5; p++ {
		go func() {
			defer wg.Done()
			totalCountLock.Lock()
			// fmt.Print("正在统计第", p, "页，")
			totalCount += 100 // totalCount = totalCount + 100 // 注意：这里有重复覆盖的问题
			// totalCountLock.Unlock()
		}() // fatal error: all goroutines are asleep - deadlock!
	}
	wg.Wait()
	fmt.Println("预计有", 100*5, "字")
	fmt.Println("总共有", totalCount, "字")
}
func countDictLockPrice() {
	fmt.Println("开始数")
	var totalCount int64 = 0
	totalCountLock := sync.Mutex{} // 变量名+Lock 清晰的知道锁了什么

	wg := sync.WaitGroup{}
	wg.Add(5)
	for p := 0; p < 5; p++ {
		go func(pageNum int) {
			defer wg.Done()
			totalCountLock.Lock()
			// fmt.Print("正在统计第", p, "页，")
			totalCount += 100 // totalCount = totalCount + 100 // 注意：这里有重复覆盖的问题
			if pageNum == 3 {
				time.Sleep(3 * time.Second)
			} // 当一个goroutine拿到锁之后，其他的goroutine必须等待锁的释放；如果有与锁无关的操作，将会大大降低性能
			totalCountLock.Unlock()
		}(p)
	}
	wg.Wait()
	fmt.Println("预计有", 100*5, "字")
	fmt.Println("总共有", totalCount, "字")
}
