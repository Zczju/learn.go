package main

import "fmt"

func main() {
	var data string
	{
		equipment := "纸带"
		if equipment == "纸带" {
			data = readFromPaper()
		} else if equipment == "磁带" {
			data = readFromMag()
		} else if equipment == "1.4软盘" {
			data = readFrom14Soft()
		} // todo else if ... else if ...
		fmt.Println(data)
	}

	{
		var equipment IOInterface = &Soft{}
		data = equipment.Read()
		fmt.Println(data)
	}
	{
		var equipment IOInterface = &Mag{}
		data = equipment.Read()
		fmt.Println(data)
	}
}

type IOInterface interface {
	Read() string
}

type Soft struct {
}

func (Soft) Read() string {
	return "啦啦啦啦啦软盘数据"
}

type Mag struct {
}

func (Mag) Read() string {
	return "滋滋滋磁带"
}

type Paper struct {
}

func (Paper) Read() string {
	return "从纸带读******00000"
}

func readFrom14Soft() string {
	return "啦啦啦啦啦软盘数据"
}

func readFromMag() string {
	return "滋滋滋磁带"
}

func readFromPaper() string {
	return "从纸带读******00000"
}
