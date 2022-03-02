package calculator

import (
	"fmt"
	gobmi "github.com/armstrongli/go-bmi"
)

//func CalcBMI(tall, weight float64) (bmi float64, err error) {
//	if tall <= 0 || weight <= 0 {
//		panic("Illegal Input Exits.")
//	}
//	return
//}

//func CalcBMI(tall, weight float64) (bmi float64, err error) {
//	if tall <= 0 || weight <= 0 {
//		return 0, errors.New("illegal input")
//	}
//	bmi, err = gobmi.BMI(weight, tall)
//	return bmi,err
//}

func CalcBMI(tall, weight float64) (bmi float64, err error) {
	if tall <= 0 || weight <= 0 {
		return 0, fmt.Errorf("illegal input")
	}
	bmi, _ = gobmi.BMI(weight, tall)
	return
}
