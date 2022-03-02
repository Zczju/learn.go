package main

import (
	"context"
	"fmt"
	"time"
)

func withDeadline() {
	now := time.Now()
	newTime := now.Add(1 * time.Second)
	ctx, _ := context.WithDeadline(context.TODO(), newTime)
	go tv(ctx)
	go mobile(ctx)
	go game(ctx)

	time.Sleep(2 * time.Second)
}
func tv(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("关电视")
			return
		default:
		}
		fmt.Println("看电视")
		time.Sleep(300 * time.Millisecond)
	}
}
func mobile(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("关手机")
			return
		default:
		}
		fmt.Println("玩手机")
		time.Sleep(300 * time.Millisecond)
	}
}
func game(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("关游戏机")
			return
		default:
		}
		fmt.Println("玩游戏")
		time.Sleep(300 * time.Millisecond)
	}
}
