package main

func fib(n uint) uint {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}

	return fib(n-1) + fib(n-2)
}

func funcDef(nums ...int) (addResult int) {
	// sum := 0  // 逻辑结构更简洁
	for _, item := range nums {
		addResult += item
	}

	return
}

func funcDef1(nums ...int) int {
	sum := 0
	for _, item := range nums {
		sum += item
	}

	return sum
}
