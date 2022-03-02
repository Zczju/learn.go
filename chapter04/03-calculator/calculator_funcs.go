package main

import "fmt"

func (c *Calculator) Add() int { // 新的拷贝，新的引用
	fmt.Printf("&c @ calculator_funcs = %p\n", &c)
	tempResult := c.left + c.right
	c.result = tempResult
	fmt.Println("调用add函数, c.result = ", c.result)
	return c.result
}
func (c Calculator) Sub() int {
	return c.left - c.right
}
func (c Calculator) Multiple() int {
	return c.left * c.right
}
func (c Calculator) Divide() int {
	return c.left / c.right
}
func (c Calculator) Reminder() int {
	return c.left % c.right
}
func (c *Calculator) setResult(result int) {
	c.result = result
}
