package _21_best_time_to_buy_and_sell_stock

import "testing"

func Test1(t *testing.T) {
	input := []int{7,1,5,3,6,4}
	output := maxProfit(input)
	if output != 5{
		t.Fail()
	}
}