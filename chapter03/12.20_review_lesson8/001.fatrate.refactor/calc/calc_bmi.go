package calculator

import (
	"fmt"
	_ "github.com/armstrongli/go-bmi"
	gobmi "github.com/armstrongli/go-bmi"
)

//func CalcBMI(tall, weight float64) (bmi float64, err error) {
//	bmi, err = gobmi.BMI(weight, tall)
//	if tall <= 0 || weight <= 0 {
//		return 0,err
//	}
//	return bmi,nil
//}

func CalcBMI(tall, weight float64) (bmi float64, err error) {
	if tall <= 0 || weight <= 0 {
		return 0, fmt.Errorf("illegal input")
	}
	bmi, _ = gobmi.BMI(weight, tall)
	return
}
