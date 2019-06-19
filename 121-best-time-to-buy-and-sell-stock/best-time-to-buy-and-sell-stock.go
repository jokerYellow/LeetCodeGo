package _21_best_time_to_buy_and_sell_stock

/*
https://leetcode.com/problems/best-time-to-buy-and-sell-stock/submissions/
121. Best Time to Buy and Sell Stock
Easy

2663

127

Favorite

Share
Say you have an array for which the ith element is the price of a given stock on day i.

If you were only permitted to complete at most one transaction (i.e., buy one and sell one share of the stock), design an algorithm to find the maximum profit.

Note that you cannot sell a stock before you buy one.

Example 1:

Input: [7,1,5,3,6,4]
Output: 5
Explanation: Buy on day 2 (price = 1) and sell on day 5 (price = 6), profit = 6-1 = 5.
             Not 7-1 = 6, as selling price needs to be larger than buying price.
Example 2:

Input: [7,6,4,3,1]
Output: 0
Explanation: In this case, no transaction is done, i.e. max profit = 0.

*/

func maxProfit(prices []int) int {
	length := len(prices)
	if length == 0 {
		return 0
	}
	maxs := make([]int, length-1)
	max := prices[length-1]
	for j := length - 2; j >= 0; j-- {
		if max < prices[j] {
			max = prices[j]
		}
		maxs[j] = max
	}
	maxValue := 0
	for i := 0; i < length-1; i++ {
		value := maxs[i] - prices[i]
		if maxValue < value {
			maxValue = value
		}
	}
	return maxValue
}
