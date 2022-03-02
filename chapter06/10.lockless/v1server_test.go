package main

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestWebServerV1(t *testing.T) {
	v1 := &WebServerV1{}
	go v1.ReloadWorker()

	start := time.Now()
	wg := sync.WaitGroup{}
	wg.Add(10000)
	for i := 0; i < 10000; i++ {
		go func() {
			defer wg.Done()
			for i := 0; i < 10000; i++ {
				v1.Visit()
			}
		}()
	}
	wg.Wait()
	finish := time.Now()
	fmt.Println("总时间:", finish.Sub(start))

}
