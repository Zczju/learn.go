package gobmi

import "fmt"

func BMI(weightKG, heightM float64) (bmi float64, error error) {
	if weightKG < 0 {
		error = fmt.Errorf("weight cannot be negative")
		return
	}
	if heightM < 0 {
		error = fmt.Errorf("height cannot be negative")
		return
	}
	if heightM == 0 {
		error = fmt.Errorf("height cannot be 0")
		return
	}
	bmi = weightKG / (heightM * heightM)
	return
}
