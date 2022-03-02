package main

import "fmt"

func main() {
	bmis := []float64{1.2, 3.22, 4.34343}
	fmt.Println(bmis)
	calculateAvg()

	sum := 1
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)

	a := 2
	for i := 0; i < 3; i++ {
		test(a)
	}
	a = 3
	test(a)

}

func calculateAvg(bmis ...float64) (Avgbmi float64) {
	var total float64
	for _, item := range bmis {
		total = total + item
	}
	return total / float64(len(bmis))
}

func test(a int) {
	a++
	fmt.Println(a)
}
