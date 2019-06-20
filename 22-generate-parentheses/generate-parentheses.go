package _2_generate_parentheses

import "fmt"

/*
https://leetcode.com/problems/generate-parentheses/
22. Generate Parentheses
Medium

2835

176

Favorite

Share
Given n pairs of parentheses, write a function to generate all combinations of well-formed parentheses.

For example, given n = 3, a solution set is:

[
  "((()))",
  "(()())",
  "(())()",
  "()(())",
  "()()()"
]
*/

func generateParenthesis(n int) []string {
	var list []string
	backtrack(&list, "", 0, 0, n)
	return list
}

func backtrack(list *[]string, str string, open, close, max int) {
	fmt.Println(list, str, open, close, max)
	if len(str) == max*2 {
		*list = append(*list, str)
		return
	}
	if open < max {
		backtrack(list, str+"(", open+1, close, max)
	}
	if close < open {
		backtrack(list, str+")", open, close+1, max)
	}
}
