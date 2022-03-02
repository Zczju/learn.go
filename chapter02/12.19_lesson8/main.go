package main

import (
	"fmt"
	"runtime/debug"
)

func main() {
	coverType()
	fmt.Println("finish")
}
func coverType() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("panic:", r)
			debug.PrintStack()
		}
	}()
	var a interface{} // any
	a = "string aaa"

	b := a.(int)
	fmt.Println(b)
}
