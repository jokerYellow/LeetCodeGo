
package _26

import (
	"testing"
)

func TestNormal826(t *testing.T) {
	difficulty := []int{4, 2, 6, 8, 10}
	profit := []int{20, 10, 30, 40, 50}
	worker := []int{4, 5, 6, 7}
	Expect := 100
	Output := maxProfitAssignment(difficulty, profit, worker)
	if Output != Expect {
		t.Fail()
	}
}

func TestNormal826_1(t *testing.T) {
	difficulty := []int{85, 47, 57}
	profit := []int{24, 66, 99}
	worker := []int{40, 25, 25}
	Expect := 0
	Output := maxProfitAssignment(difficulty, profit, worker)
	if Output != Expect {
		t.Fail()
	}
}

func TestNormal826_2(t *testing.T) {
	difficulty := []int{68, 35, 52, 47, 86}
	profit := []int{67, 17, 1, 81, 3}
	worker := []int{92, 10, 85, 84, 82}
	Expect := 324
	Output := maxProfitAssignment(difficulty, profit, worker)
	if Output != Expect {
		t.Fail()
	}
}
