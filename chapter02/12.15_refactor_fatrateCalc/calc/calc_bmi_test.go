package calculator

import "testing"

func TestCalcBMI(t *testing.T) {
	inputHeight, inputWeight := 1.0, 1.0
	expectOutput := 1.0
	t.Logf("开始计算，输入：height：%f,weight: %f,期望结果：%f", inputWeight, inputWeight, expectOutput)
	actualOutput, err := CalcBMI(inputHeight, inputWeight)
	if err != nil {
		t.Fatalf("expecting no error, but got : %v", err)
	}
	if expectOutput != actualOutput {
		t.Errorf("expecting %f,but got %f", expectOutput, actualOutput)
	}
}
