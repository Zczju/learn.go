package main

import (
	"fmt"
	"sync"
	"time"
)

type Store struct {
	DataCount int
	Max       int
	lock      sync.Mutex // 这里Mutex是一个对象，Store初始化的时候会自动初始化；如果是引用类型，则要手动初始化
}

type Producer struct {
}

func (Producer) Produce(s *Store) {
	s.lock.Lock()
	defer s.lock.Unlock()
	if s.DataCount == s.Max {
		fmt.Println("生产者看到库存满了，不生产")
		return
	}
	fmt.Println("开始生产 库存+1")
	s.DataCount++
}

type Consumer struct {
}

func (Consumer) Consume(s *Store) {
	s.lock.Lock()
	defer s.lock.Unlock()
	if s.DataCount == 0 {
		fmt.Println("消费者看到没库存了，不消费")
		return
	}
	fmt.Println("开始消费 库存-1")
	s.DataCount--
}

func main() {
	s := &Store{
		Max: 10,
	}
	pCount, cCount := 50, 50
	for i := 0; i < pCount; i++ {
		go func() {
			for {
				time.Sleep(100 * time.Millisecond)
				Producer{}.Produce(s)
			}
		}()
	}
	for i := 0; i < cCount; i++ {
		go func() {
			for {
				time.Sleep(100 * time.Millisecond)
				Consumer{}.Consume(s)
			}

		}()
	}
	time.Sleep(1 * time.Second) // 保证所有的go routine能够正常运行完，延长主进程的退出时间
	fmt.Println(s.DataCount)
}
