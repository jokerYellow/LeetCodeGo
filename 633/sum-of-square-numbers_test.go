package leetcode

import (
	"math"
	"testing"
)

func TestN1(t *testing.T) {
	a := 3
	if judgeSquareSum(a) == true {
		t.Fail()
	}
}

func TestN2(t *testing.T) {
	a := 25000*25000 + 1000*1000
	if judgeSquareSum(a) == false {
		t.Fail()
	}
}

func TestN3(t *testing.T) {
	a := 999999999
	if judgeSquareSum(a) == true {
		t.Fail()
	}
}

func TestN4(t *testing.T) {
	a := math.MaxInt16
	if judgeSquareSum(a) == true {
		t.Fail()
	}
}
