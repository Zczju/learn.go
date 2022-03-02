package main

import (
	"fmt"
	"sync"
	"time"
)

type Store struct {
	DataCount int
	Max       int
	lock      sync.Mutex // 这里Mutex是一个对象，Store初始化的时候会自动初始化；如果是引用类型，则要手动初始化，比如下面两个cond
	pCond     *sync.Cond
	cCond     *sync.Cond
}

type Producer struct {
}

func (p Producer) Produce(s *Store) {
	s.lock.Lock()
	defer s.lock.Unlock()
	if s.DataCount == s.Max {
		fmt.Println("生产者在等货")
		s.pCond.Wait()
	}
	fmt.Println("开始生产 库存+1")
	s.DataCount++
	s.cCond.Signal()
}

type Consumer struct {
}

func (c Consumer) Consume(s *Store) {
	s.lock.Lock()
	defer s.lock.Unlock()
	if s.DataCount == 0 {
		fmt.Println("消费者在等清仓")
		s.cCond.Wait()
	}
	fmt.Println("开始消费 库存-1")
	s.DataCount--
	s.pCond.Signal()
}

func main() {
	s := &Store{
		Max: 10,
	}
	s.pCond = sync.NewCond(&s.lock) // 通过系统包来初始化，不能直接在对象中初始化
	s.cCond = sync.NewCond(&s.lock) // 这个实例化有风险，因为都等在同一把锁下

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
