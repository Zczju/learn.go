package main

import (
	"context"
	"fmt"
	"time"
)

func withCancel() {
	ctx := context.TODO()
	ctx, cancel := context.WithCancel(ctx)
	fmt.Println("做蛋挞，要买材料")
	go buyFlour(ctx)
	go buyOil(ctx)
	go buyEgg(ctx)
	time.Sleep(500 * time.Millisecond)
	fmt.Println("没电了，取消购买所有食材")
	cancel() // 当调用cancel后，所有由此上下文衍生出的context都会cancel
	time.Sleep(1 * time.Second)
}
func buyFlour(ctx context.Context) {
	fmt.Println("去买面")
	time.Sleep(1 * time.Second)

	select {
	case <-ctx.Done():
		fmt.Println("收到消息，不买面了")
		return
	default:
	}
	fmt.Println("买面")
}
func buyOil(ctx context.Context) {
	fmt.Println("去买油")
	time.Sleep(1 * time.Second)

	select {
	case <-ctx.Done():
		fmt.Println("收到消息，不买油了")
		return
	default:
	}
	fmt.Println("买油")
}
func buyEgg(ctx1 context.Context) {
	ctx, _ := context.WithCancel(ctx1)
	// defer cancel() // 无论是否调用衍生出来的ctx的cancel，Done返回的channel都会关闭。
	fmt.Println("去买蛋")
	// time.Sleep(1 * time.Second)

	select {
	case <-ctx.Done():
		fmt.Println("收到消息，不买蛋了")
		return
	default:
	}
	fmt.Println("买蛋")
	go buySEgg(ctx)
	go buyBEgg(ctx)
	time.Sleep(1 * time.Second)
}
func buySEgg(ctx context.Context) {
	fmt.Println("去买蛋: buySEgg")
	time.Sleep(1 * time.Second)

	select {
	case <-ctx.Done():
		fmt.Println("收到消息，不买蛋了: buySEgg")
		return
	default:
	}
	fmt.Println("买蛋")
}
func buyBEgg(ctx context.Context) {
	fmt.Println("去买蛋: buyBEgg")
	time.Sleep(1 * time.Second)

	select {
	case <-ctx.Done():
		fmt.Println("收到消息，不买蛋了: buyBEgg")
		return
	default:
	}
	fmt.Println("买蛋")
}
