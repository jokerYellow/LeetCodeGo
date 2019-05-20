package _18_the_skyline_problem

import (
	"fmt"
	"github.com/jokerYellow/leetcode/utils"
	"testing"
)

func Test1(t *testing.T) {
	intput := [][]int{
		{2,9,10},
		{3,7,15},
		{5,12,12},
		{15,20,10},
		{19,24,8},
	}
	output := getSkyline(intput)
	expect := [][]int{
		{2,10},
		{3,15},
		{7,12},
		{12,0},
		{15,10},
		{20,8},
		{24,0},
	}
	fmt.Println("expect:",expect)
	fmt.Println("output:",output)
	if utils.CheckEqual(output,expect) == false {
		t.Fail()
	}
}