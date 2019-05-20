package _18_the_skyline_problem

import (
	"fmt"
	"github.com/jokerYellow/leetcode/utils"
	"testing"
)

func Test1(t *testing.T) {
	intput := [][]int{
		{2, 9, 10},
		{3, 7, 15},
		{5, 12, 12},
		{15, 20, 10},
		{19, 24, 8},
	}
	output := getSkyline(intput)
	expect := [][]int{
		{2, 10},
		{3, 15},
		{7, 12},
		{12, 0},
		{15, 10},
		{20, 8},
		{24, 0},
	}
	fmt.Println("expect:", expect)
	fmt.Println("output:", output)
	if utils.CheckEqual(output, expect) == false {
		t.Fail()
	}
}

func Test2(t *testing.T) {

	intput := [][]int{
		{2, 9, 20},
		{3, 7, 15},
		{5, 12, 1},
		{15, 100, 1},
		{19, 24, 8},
	}
	output := getSkyline(intput)
	expect := [][]int{
		{2, 20},
		{9, 1},
		{12, 0},
		{15, 1},
		{19, 8},
		{24, 1},
		{100, 0},
	}
	fmt.Println("expect:", expect)
	fmt.Println("output:", output)
	if utils.CheckEqual(output, expect) == false {
		t.Fail()
	}
}

func Test3(t *testing.T) {

	intput := [][]int{
		{0, 2, 3},
		{2, 5, 3},
	}
	output := getSkyline(intput)
	expect := [][]int{
		{0, 3},
		{5, 0},
	}
	fmt.Println("expect:", expect)
	fmt.Println("output:", output)
	if utils.CheckEqual(output, expect) == false {
		t.Fail()
	}
}

func Test4(t *testing.T) {

	intput := [][]int{
		{1, 2, 1},
		{1, 2, 2},
		{1, 2, 3},
	}
	output := getSkyline(intput)
	expect := [][]int{
		{1, 3},
		{2, 0},
	}
	fmt.Println("expect:", expect)
	fmt.Println("output:", output)
	if utils.CheckEqual(output, expect) == false {
		t.Fail()
	}
}

func Test5(t *testing.T) {

	intput := [][]int{
		{0, 3, 3},
		{1, 5, 3},
		{2, 4, 3},
		{3, 7, 3},
	}
	output := getSkyline(intput)
	expect := [][]int{
		{0, 3},
		{7, 0},
	}
	fmt.Println("expect:", expect)
	fmt.Println("output:", output)
	if utils.CheckEqual(output, expect) == false {
		t.Fail()
	}
}

func Test6(t *testing.T) {

	intput := [][]int{
		{1, 5, 3}, {1, 5, 3}, {1, 5, 3},
	}
	output := getSkyline(intput)
	expect := [][]int{
		{1, 3},
		{5, 0},
	}
	fmt.Println("expect:", expect)
	fmt.Println("output:", output)
	if utils.CheckEqual(output, expect) == false {
		t.Fail()
	}
}
