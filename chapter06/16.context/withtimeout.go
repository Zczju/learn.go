package main

import (
	"context"
	"fmt"
	"time"
)

func withTimeout() {
	ctx, _ := context.WithTimeout(context.TODO(), 1*time.Second)
	fmt.Println("开始部署望远镜，发送信号")
	go distributeMainFrame(ctx)
	go distributeMainBody(ctx)
	go distributeCover(ctx)
	select {
	case <-ctx.Done():
		fmt.Println("任务超时没有完成")
	}
	time.Sleep(11 * time.Second) // 等待11秒后收到任务取消的消息

}

func distributeMainFrame(ctx context.Context) {
	time.Sleep(10 * time.Second)
	select {
	case <-ctx.Done():
		fmt.Println("任务失败，停止部署: distributeMainFrame")
		return
	default:
	}
	fmt.Println("部署: distributeMainFrame")
}
func distributeMainBody(ctx context.Context) {
	time.Sleep(10 * time.Second)
	select {
	case <-ctx.Done():
		fmt.Println("任务失败，停止部署: distributeMainBody")
		return
	default:
	}
	fmt.Println("部署: distributeMainBody")
}
func distributeCover(ctx context.Context) {
	time.Sleep(10 * time.Second)
	select {
	case <-ctx.Done():
		fmt.Println("任务失败，停止部署: distributeCover")
		return
	default:
	}
	fmt.Println("部署: distributeCover")
}
