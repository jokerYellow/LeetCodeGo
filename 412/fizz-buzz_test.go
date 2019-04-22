package leetcode

import (
	"github.com/jokerYellow/leetcode/utils"
	"testing"
)

func Test1(t *testing.T) {
	input := 15
	output := fizzBuzz(input)
	expect := []string{"1",
		"2",
		"Fizz",
		"4",
		"Buzz",
		"Fizz",
		"7",
		"8",
		"Fizz",
		"Buzz",
		"11",
		"Fizz",
		"13",
		"14",
		"FizzBuzz"}
	if !utils.EqualStrings(output,expect){
		t.Fail()
	}
}