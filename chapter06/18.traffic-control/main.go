package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	for i := 0; i < 2; i++ {
		go sendRequest()
	}
	time.Sleep(5 * time.Second)

}

func sendRequest() {
	for {
		rand.Seed(time.Now().UnixNano())
		time.Sleep(time.Duration(rand.Int63()%30) * time.Millisecond)
		serveRequest()
	}
}
func serveRequest() {
	accept := trafficControlStart()
	if accept {
		defer trafficControlFinish()
		fmt.Println("服务请求")
		rand.Seed(time.Now().UnixNano())
		time.Sleep(time.Duration(rand.Int63()%30) * time.Millisecond)
		// 模拟服务请求的时间
	} else {
		fmt.Println("服务请求被拒绝")
	}

}

var trafficControlCh = make(chan struct{}, 1) // 无法做动态变更

func trafficControlStart() (accept bool) {
	select {
	case trafficControlCh <- struct{}{}:
		fmt.Println("接受请求")
		return true
	default:
		fmt.Println("拒绝请求")
		return
	}

}
func trafficControlFinish() {
	<-trafficControlCh
}
