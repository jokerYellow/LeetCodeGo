package _07_trapping_rain_water_ii

import (
	"fmt"
	"testing"
)

func Test1(t *testing.T) {
	input := [][]int{
		{1, 4, 3, 1, 3, 2},
		{3, 2, 1, 3, 2, 4},
		{2, 3, 3, 2, 3, 1},
	}
	expect := 4
	output := trapRainWater(input)
	fmt.Println(input)
	fmt.Println(expect)
	fmt.Println(output)
	if expect != output {
		t.Fail()
	}
}

func Test2(t *testing.T) {
	input := [][]int{
		{10, 10, 0, 10, 10, 10},
		{10, 0, 0, 0, 0, 10},
		{10, 0, 0, 0, 0, 10},
		{10, 10, 10, 10, 10, 10},
	}
	expect := 0
	output := trapRainWater(input)
	fmt.Println(input)
	fmt.Println(expect)
	fmt.Println(output)
	if expect != output {
		t.Fail()
	}
}

func Test3(t *testing.T) {
	input := [][]int{
		{1, 4, 3, 1, 3, 2},
		{3, 2, 1, 3, 2, 4},
		{2, 3, 3, 2, 3, 1},
	}
	expect := 4
	output := trapRainWater(input)
	fmt.Println(input)
	fmt.Println(expect)
	fmt.Println(output)
	if expect != output {
		t.Fail()
	}
}

func Test4(t *testing.T) {
	input := [][]int{
		{1, 4, 3, 1, 3, 2},
		{3, 2, 1, 3, 2, 4},
		{2, 3, 3, 2, 3, 1},
	}
	expect := 4
	output := trapRainWater(input)
	fmt.Println(input)
	fmt.Println(expect)
	fmt.Println(output)
	if expect != output {
		t.Fail()
	}
}
