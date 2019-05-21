package _3_roman_to_integer

import "testing"

func Test1(t *testing.T) {
	input := "III"
	expect := 3
	output := romanToInt(input)
	if expect != output{
		t.Fail()
	}
}

func Test2(t *testing.T) {
	input := "VI"
	expect := 6
	output := romanToInt(input)
	if expect != output{
		t.Fail()
	}
}

func Test3(t *testing.T) {
	input := "IV"
	expect := 4
	output := romanToInt(input)
	if expect != output{
		t.Fail()
	}
}
func Test4(t *testing.T) {
	input := "LVIII"
	expect := 58
	output := romanToInt(input)
	if expect != output{
		t.Fail()
	}
}

func Test5(t *testing.T) {
	input := "MCMXCIV"
	expect := 1994
	output := romanToInt(input)
	if expect != output{
		t.Fail()
	}
}

func Test6(t *testing.T) {
	input := "MMMCCC"
	expect := 3300
	output := romanToInt(input)
	if expect != output{
		t.Fail()
	}
}
