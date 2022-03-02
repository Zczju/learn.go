package main

import (
	"fmt"
	"log"
	"os"
)

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	//file, err := os.OpenFile("aardvark.txt", os.O_APPEND,os.FileMode(0600))
	//check(err)
	//defer file.Close()
	//_, err = file.Write([]byte("amazing!\n"))
	//check(err)
	//err = file.Close()
	//check(err)
	//scanner := bufio.NewScanner(file)
	//for scanner.Scan(){
	//	fmt.Println(scanner.Text())
	//}
	//check(scanner.Err())

	//fmt.Printf("%3d: %08b\n", 0, 0)
	//fmt.Printf("%3d: %08b\n", 1, 1)
	//fmt.Printf("%3d: %08b\n", 2, 2)
	//fmt.Printf("%3d: %08b\n", 3, 3)
	//fmt.Printf("%3d: %08b\n", 4, 4)
	//fmt.Printf("%3d: %08b\n", 5, 5)
	//fmt.Printf("%3d: %08b\n", 6, 6)
	//fmt.Printf("%3d: %08b\n", 7, 7)
	//fmt.Printf("%3d: %08b\n", 8, 8)
	//fmt.Printf("%3d: %08b\n", 9, 9)
	//fmt.Printf("%3d: %08b\n", 16, 16)
	//fmt.Printf("%3d: %08b\n", 32, 32)
	//fmt.Printf("%3d: %08b\n", 128, 128)

	//fmt.Println(os.O_RDONLY, os.O_WRONLY, os.O_RDWR, os.O_CREATE, os.O_APPEND)
	//fmt.Printf("%016b\n", os.O_RDONLY)
	//fmt.Printf("%016b\n", os.O_WRONLY)
	//fmt.Printf("%016b\n", os.O_RDWR)
	//fmt.Printf("%016b\n", os.O_CREATE)
	//fmt.Printf("%016b\n", os.O_APPEND)
	//fmt.Printf("%016b\n", os.O_WRONLY|os.O_CREATE)
	//fmt.Printf("%016b\n", os.O_WRONLY|os.O_CREATE|os.O_APPEND)

	//option := os.O_WRONLY|os.O_APPEND|os.O_CREATE
	//file, err := os.OpenFile("aardvark.txt",option ,os.FileMode(0600))
	//check(err)
	//_, err = file.Write([]byte("amazing!\n"))
	//check(err)
	//err = file.Close()
	//check(err)

	fmt.Println(os.FileMode(0700))
	for i := 0; i < 19; i++ {
		fmt.Printf("%3d: %04o\n", i, i)
	}
	fmt.Println(0007)
}
