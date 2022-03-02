package main

import "fmt"

var weight1 float64 = 1.345

func main() {
	changeWeight(2.345)
	fmt.Println(weight1)
}

func changeWeight(weight2 float64) {
	weight1 += weight2
}
