package main

import (
	"fmt"
	"reflect"
)

func main() {
	var refrigerator Refrigerator
	fmt.Println(refrigerator.Size)
	var elephant Elephant
	fmt.Println(elephant.Name)

	// var ip *int // 默认是nil
	// fmt.Println(*ip) // panic: 空指针
	// m := map[string]string{} // map可以直接初始化实体，但是接口不行

	// var putER PutElephantIntoRefrigerator // 默认是nil
	// putER.OpenTheDoorOfRefrigerator(refrigerator) // panic: runtime error: invalid memory address or nil pointer dereference

	// **注意：如果某个对象的成员方法有的在对象上，有的在对象指针上，那么在给接口赋值时，必须用指针。**
	var legend PutElephantIntoRefrigerator = &PutElephantIntoRefrigeratorImpl{}
	legend.OpenTheDoorOfRefrigerator(refrigerator)
	legend.PutElephantIntoRefrigerator(elephant, refrigerator)
	legend.CloseTheDoorOfRefrigerator(refrigerator)
	// todo show the elephant in refrigerator

	// var o Open = Refrigerator{}
	var c Close = Refrigerator{} // 小的，只有一个接口
	var b Box = Refrigerator{}   // 大的，有两个接口
	fmt.Println(b, c)
	c = b // 用大的给小的赋值，可行
	// b = c // 用小的给大的赋值，不可行

	var i interface{}
	i = 3
	fmt.Println(reflect.TypeOf(i), "value:", i)
}

type PutElephantIntoRefrigerator interface {
	OpenTheDoorOfRefrigerator(Refrigerator) error
	PutElephantIntoRefrigerator(Elephant, Refrigerator) error
	CloseTheDoorOfRefrigerator(Refrigerator) error
}

type Refrigerator struct {
	Size string
}

func (r Refrigerator) Open() error {
	return nil
}

func (r Refrigerator) Close() error {
	return nil
}

type Elephant struct {
	Name string
}

type TestTypeImplInterface func()

func (t TestTypeImplInterface) OpenTheDoorOfRefrigerator(Refrigerator) error {
	return nil
}

func (t TestTypeImplInterface) PutElephantIntoRefrigerator(Elephant, Refrigerator) error {
	return nil
}

func (t TestTypeImplInterface) CloseTheDoorOfRefrigerator(Refrigerator) error {
	return nil
}

type PutElephantIntoRefrigeratorImpl struct {
}

func (legent *PutElephantIntoRefrigeratorImpl) OpenTheDoorOfRefrigerator(refrigerator Refrigerator) error {
	// todo
	fmt.Println("打开冰箱门")
	return nil
}

func (legent PutElephantIntoRefrigeratorImpl) PutElephantIntoRefrigerator(elephant Elephant, refrigerator Refrigerator) error {
	// todo
	fmt.Println("装进去")
	return nil
}

func (legent *PutElephantIntoRefrigeratorImpl) CloseTheDoorOfRefrigerator(refrigerator Refrigerator) error {
	// todo
	fmt.Println("关门")
	return nil
}

// * 必须实现所有的接口方法，才叫做实现接口，否则不算; 多了方法没事
// * 接口的实现，不限类型。 对象可以为任意类型，只要成员方法全部实现
// * 成员函数可以在成员对象上实现，也可以在对象指针上实现。在给接口变量赋值时，如果是混合定义(有定义在对象上的也有定义在指针上的)，要用指针

type Open interface {
	Open() error
}

type Close interface {
	Close() error
}

type Box interface {
	Open
	Close
}
