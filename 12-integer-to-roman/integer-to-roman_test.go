package _2_integer_to_roman

import "testing"

func Test1(t *testing.T) {
	input := 3
	expect := "III"
	output := intToRoman(input)
	if expect != output {
		t.Fail()
	}
}

func Test2(t *testing.T) {
	input := 4
	expect := "IV"
	output := intToRoman(input)
	if expect != output {
		t.Fail()
	}
}

func Test3(t *testing.T) {
	input := 9
	expect := "IX"
	output := intToRoman(input)
	if expect != output {
		t.Fail()
	}
}

func Test4(t *testing.T) {
	input := 58
	expect := "LVIII"
	output := intToRoman(input)
	if expect != output {
		t.Fail()
	}
}

func Test5(t *testing.T) {
	input := 1994
	expect := "MCMXCIV"
	output := intToRoman(input)
	if expect != output {
		t.Fail()
	}
}

func Test6(t *testing.T) {
	input := 3000
	expect := "MMM"
	output := intToRoman(input)
	if expect != output {
		t.Fail()
	}
}
