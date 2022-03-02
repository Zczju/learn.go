package main

import (
	"fmt"
	"sync"
)

type Queue struct {
	data []interface{}
	sync.Mutex
}

func (q *Queue) Push(data interface{}) {
	q.Lock()
	defer q.Unlock()
	q.data = append(q.data, data) // append方法将数据加到切片末尾
}
func (q *Queue) Pop() (interface{}, bool) {
	q.Lock()
	defer q.Unlock()
	// 模拟队列为先进先出，数据从尾部加入，从头部输出
	if len(q.data) > 0 {
		o := q.data[0]
		q.data = q.data[1:]
		return o, true
	}
	return nil, false
}

func main() {
	q := &Queue{data: []interface{}{}}
	q.Push(111)
	q.Push(222)
	q.Push(333)
	q.Push(nil)

	fmt.Println(q.Pop()) // 111
	fmt.Println(q.Pop()) // 222
	fmt.Println(q.Pop()) // 333
	fmt.Println(q.Pop()) // nil true
	fmt.Println(q.Pop()) // nil false

}
