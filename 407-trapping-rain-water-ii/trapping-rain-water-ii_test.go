package _07_trapping_rain_water_ii

import (
	"fmt"
	"testing"
)

func TestHeap(t *testing.T) {
	h := &heap{}
	h.items = []int{1, 2, 3, 4, 3, 2, 1, 2, 2}
	fmt.Println(h)
	h.maxHeapify(1)
	fmt.Println(h)
}

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
