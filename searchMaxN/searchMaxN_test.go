package searchMaxN

import "testing"

func TestSearch(t *testing.T) {
	testSearch(t, []int{1, 2, 3, 4}, 3, 2)
	testSearch(t, []int{1, 2, 3, 4, 5, 6, 7}, 6, 2)
}

func testSearch(t *testing.T, input []int, expect, index int) {
	output := searchMaxN(input, index)
	if output != expect {
		t.Fail()
	}
}
