package __reverse_integer

import "testing"

func Test1(t *testing.T) {
	input := 1000004
	output := reverse(input)
	if output != 4000001{
		t.Fail()
	}
}


func Test2(t *testing.T) {
	input := -1000004
	output := reverse(input)
	if output != -4000001{
		t.Fail()
	}
}