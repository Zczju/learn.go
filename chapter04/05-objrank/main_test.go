package main

import "testing"

func TestCase1Part1(t *testing.T) {
	r := &FatRateRank{}
	r.inputRecord("王强", 0.38)
	r.inputRecord("王强", 0.32)
	{
		rankOfWQ, fatRateOfWQ := r.getRank("王强")
		if rankOfWQ != 1 {
			t.Fatalf("WQ期望排名为第1，但结果却为 %d", rankOfWQ)
		}
		if fatRateOfWQ != 0.32 {
			t.Fatalf("WQ期望体脂率为0.32，但结果却为 %f", fatRateOfWQ)
		}
	}
}

func TestCase1(t *testing.T) {
	r := &FatRateRank{}
	r.inputRecord("王强", 0.38)
	r.inputRecord("王强", 0.32)
	{
		rankOfWQ, fatRateOfWQ := r.getRank("王强")
		if rankOfWQ != 1 {
			t.Fatalf("WQ期望排名为第1，但结果却为 %d", rankOfWQ)
		}
		if fatRateOfWQ != 0.32 {
			t.Fatalf("WQ期望体脂率为0.32，但结果却为 %f", fatRateOfWQ)
		}
	}

	r.inputRecord("李静", 0.28)
	rankOfWQ, fatRateOfWQ := r.getRank("王强")
	rankOfLJ, fatRateOfLJ := r.getRank("李静")
	if rankOfWQ != 2 {
		t.Fatalf("WQ期望排名为第2，但结果却为 %d", rankOfWQ)
	}
	if fatRateOfWQ != 0.32 {
		t.Fatalf("WQ期望体脂率为0.32，但结果却为 %f", fatRateOfWQ)
	}

	{
		if rankOfLJ != 1 {
			t.Fatalf("LJ期望排名为第1，但结果却为 %d", rankOfLJ)
		}
		if fatRateOfLJ != 0.28 {
			t.Fatalf("LJ期望体脂率为0.28，但结果却为 %f", fatRateOfLJ)
		}
	}
}

func TestCase2(t *testing.T) {
	r := &FatRateRank{}
	r.inputRecord("王强", 0.38)
	r.inputRecord("张伟", 0.38)
	r.inputRecord("李静", 0.28)
	{
		rankOfWQ, fatRateOfWQ := r.getRank("王强")
		if rankOfWQ != 2 {
			t.Fatalf("WQ期望排名为第2，但结果却为 %d", rankOfWQ)
		}
		if fatRateOfWQ != 0.38 {
			t.Fatalf("WQ期望体脂率为0.38，但结果却为 %f", fatRateOfWQ)
		}
	}
	{
		rankOfLJ, fatRateOfLJ := r.getRank("李静")
		if rankOfLJ != 1 {
			t.Fatalf("LJ期望排名为第1，但结果却为 %d", rankOfLJ)
		}
		if fatRateOfLJ != 0.28 {
			t.Fatalf("LJ期望体脂率为0.28，但结果却为 %f", fatRateOfLJ)
		}
	}
	{
		rankOfZW, fatRateOfZW := r.getRank("张伟")
		if rankOfZW != 2 {
			t.Fatalf("ZW期望排名为第2，但结果却为 %d", rankOfZW)
		}
		if fatRateOfZW != 0.38 {
			t.Fatalf("ZW期望体脂率为0.38，但结果却为 %f", fatRateOfZW)
		}
	}
}

func TestCase3(t *testing.T) {
	r := &FatRateRank{}
	r.inputRecord("王强", 0.38)
	r.inputRecord("张伟")
	r.inputRecord("李静", 0.28)
	{
		rankOfWQ, fatRateOfWQ := r.getRank("王强")
		if rankOfWQ != 2 {
			t.Fatalf("WQ期望排名为第2，但结果却为 %d", rankOfWQ)
		}
		if fatRateOfWQ != 0.38 {
			t.Fatalf("WQ期望体脂率为0.38，但结果却为 %f", fatRateOfWQ)
		}
	}
	{
		rankOfLJ, fatRateOfLJ := r.getRank("李静")
		if rankOfLJ != 1 {
			t.Fatalf("LJ期望排名为第1，但结果却为 %d", rankOfLJ)
		}
		if fatRateOfLJ != 0.28 {
			t.Fatalf("LJ期望体脂率为0.28，但结果却为 %f", fatRateOfLJ)
		}
	}
	{
		rankOfZW, _ := r.getRank("张伟")
		if rankOfZW != 3 {
			t.Fatalf("ZW期望排名为第3，但结果却为 %d", rankOfZW)
		}
	}
}
