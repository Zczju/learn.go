package main

import "fmt"

func main() {
	a, b := 1, 2
	add(&a, &b)
	fmt.Println(a)

	c := &a
	d := &c
	fmt.Println("d=", d, "*d=", *d, "**d=", **d)
	fmt.Println("c= ", c, "*c= ", *c)

	m := map[string]string{} // map和切片都是引用类型,指针也是一个引用类型
	mp1 := &m                //map pointer 1 的类型是 *map[string]string
	fmt.Println(*mp1)
	//put(*mp1) // 两种输入方式结果相同
	put(m) // 此处传的不是指针
	fmt.Println("*mp1= ", *mp1)

	f1 := add // f1 = func (*int, *int)
	f1(&a, &b)
	fmt.Println("f1, add= ", a)
	f1p1 := &f1 // f1p1 = *func (*int, *int)
	(*f1p1)(&a, &b)
	// 要继续使用必须先拿到函数本身，拿*打开他的盒子，再去调用
	fmt.Println("f1p1, add= ", a)

	{
		var nothing *int // 不存在的盒子无法装东西
		// *nothing = 3 // 注意: 这里是没有指向任何东西的int指针
		//panic: runtime error: invalid memory address or nil pointer dereference
		fmt.Println("nothing= ", nothing) // nil 空指针
	}
	{
		var nothingMap map[string]string
		// nothingMap["a"] = "BBB" // map必须实例化才能写入
		// panic: assignment to entry in nil map
		fmt.Println(nothingMap)
	}
	{
		var nothingSlice []int
		// nothingSlice[0] = 100
		// panic: runtime error: index out of range [0] with length 0
		nothingSlice = append(nothingSlice, 100) // append可以使用，将元素加进去（切片的特殊性）
		fmt.Println(nothingSlice)
	}
}

func add(a, b *int) {
	*a = *a + *b
}

func put(m map[string]string) {
	m["a"] = "AAA"
}
