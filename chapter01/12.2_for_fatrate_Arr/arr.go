package main

import (
	"fmt"
	"reflect"
)

func main() {
	var ages1 [5]int = [5]int{35, 34, 54, 65, 33}
	fmt.Println(ages1)
	var ages2 = [5]int{1, 2, 3, 4, 5}
	ages1 = ages2
	fmt.Println(ages1)

	ages1 = reverse(ages1)
	fmt.Println(ages1)

	fmt.Println("3rd" +
		" version")
	//难点 数组长度管理
	newPersonInfos := [...][3]string{
		{},
		{},
		{},
		{},
	}
	for _, val := range newPersonInfos {
		fmt.Println(val)
	}

	//用降维方式
	for d1, d1val := range newPersonInfos {
		for d2, d2val := range d1val {
			fmt.Println(d1, d2, d2val)
		}
	}

	a := [6]int{1, 2, 3, 4, 5, 6}
	Revise2(a)
	fmt.Println(a)

	slice2 := make([]string, 3, 5)
	fmt.Println(slice2)
	fmt.Printf("len:%d", len(slice2))
	fmt.Printf("cap:%d\n", cap(slice2))

	i := new(bool)
	fmt.Println(reflect.TypeOf(i))
	fmt.Println(*i)

	k := 0
	j := &k
	fmt.Println(reflect.TypeOf(j))

	//copy(slice2,ages2)

}

func Revise2(a [6]int) {
	for i := 0; i < len(a)/2; i++ {
		a[i], a[len(a)-1-i] = a[len(a)-1-i], a[i]
	}
}

func reverse(arr [5]int) [5]int {
	i, j := 0, len(arr)-1
	for i < j {
		arr[i], arr[j] = arr[j], arr[i]
		i++
		j--
	}
	return arr
}
