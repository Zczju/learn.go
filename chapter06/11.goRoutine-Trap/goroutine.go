package main

import (
	"fmt"
	"sync"
)

func main() {
	intArr := []int{1, 2, 3, 4, 5, 6, 7}
	wg := sync.WaitGroup{}
	wg.Add(len(intArr))
	for _, item := range intArr {
		go func(newItem int) {
			defer wg.Done()
			for i := 0; i < 10; i++ {
				fmt.Println(newItem)
				// 闭包导致匿名函数传入的共享变量的共享出现问题
				// 解决方法：剪断内外关联关系，将item显式的通过参数传进来
				// go routine 与 main中的for循环一起走，不确定会读到哪一个
			}
		}(item)
	}
	wg.Wait()
}
