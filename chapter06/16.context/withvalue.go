package main

import (
	"context"
	"fmt"
	"time"
)

func withValue() {
	ctx := context.WithValue(context.TODO(), "1", "钱包")
	go func(ctx context.Context) {
		time.Sleep(1 * time.Second)
		fmt.Println("withValue: 1", ctx.Value("1"))
		fmt.Println("withValue: 2", ctx.Value("2"))
		fmt.Println("withValue: 3", ctx.Value("3"))
		fmt.Println("withValue: 4", ctx.Value("4"))
	}(ctx) // 不论是新起一个goroutine，还是sleep一段时间，都不会影响parent context
	goToPapa(ctx)
	time.Sleep(2 * time.Second)
}

func goToPapa(ctx context.Context) {
	ctx = context.WithValue(ctx, "2", "充电宝")
	go func(ctx context.Context) {
		time.Sleep(1 * time.Second)
		fmt.Println("goToPapa: 1", ctx.Value("1"))
		fmt.Println("goToPapa: 2", ctx.Value("2"))
		fmt.Println("goToPapa: 3", ctx.Value("3"))
		fmt.Println("goToPapa: 4", ctx.Value("4"))
	}(ctx)
	goToMama(ctx)
}
func goToMama(ctx context.Context) {
	ctx = context.WithValue(ctx, "3", "小夹克")
	go func(ctx context.Context) {
		time.Sleep(1 * time.Second)
		fmt.Println("goToMama: 1", ctx.Value("1"))
		fmt.Println("goToMama: 2", ctx.Value("2"))
		fmt.Println("goToMama: 3", ctx.Value("3"))
		fmt.Println("goToMama: 4", ctx.Value("4"))
	}(ctx)
	goToGrandma(ctx)
}
func goToGrandma(ctx context.Context) {
	ctx = context.WithValue(ctx, "4", "大苹果")
	go func(ctx context.Context) {
		time.Sleep(1 * time.Second)
		fmt.Println("goToGrandma: 1", ctx.Value("1"))
		fmt.Println("goToGrandma: 2", ctx.Value("2"))
		fmt.Println("goToGrandma: 3", ctx.Value("3"))
		fmt.Println("goToGrandma: 4", ctx.Value("4"))
	}(ctx)
	goToParty(ctx)

}
func goToParty(ctx context.Context) {
	fmt.Println("goToParty: 1", ctx.Value("1"))
	fmt.Println("goToParty: 2", ctx.Value("2"))
	fmt.Println("goToParty: 3", ctx.Value("3"))
	fmt.Println("goToParty: 4", ctx.Value("4"))
}
