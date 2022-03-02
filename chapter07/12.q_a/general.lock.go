package main

import "sync"

var Counter *safeCount = &safeCount{
	totalNumber:      0,
	totalLetterCount: 0,
	totalWordCount:   0,
	Mutex:            sync.Mutex{},
}

type safeCount struct {
	totalNumber      int
	totalLetterCount int
	totalWordCount   int

	// .....
	sync.Mutex
}

// 做一个单独的obj，把里面的内容作为整体。一把锁去锁一堆数据，很容易写乱

// AddNumber 定义成员方法 只需要知道这是一个线程安全的AddNumber，其他不需要管，直接操作
func (c *safeCount) AddNumber(totalNumber int, totalLetterCount int, totalWordCount int) {
	c.Lock()
	defer c.Unlock()
	c.totalNumber += totalNumber
	// .....
}

func countNumber() {
	// 在此处countNumber的时候，是去调用c的成员方法去做，而不是直接操作锁
	// 操作锁的过程由AddNumber函数来完成
	Counter.AddNumber(100, 500, 10)

}
