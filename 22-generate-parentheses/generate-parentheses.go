package _2_generate_parentheses

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
	appendParenthesis(&list, "", 0, 0, n)
	return list
}

func appendParenthesis(list *[]string, text string, open, close, max int) {
	if len(text) == 2*max {
		*list = append(*list, text)
		return
	}
	if open < max {
		appendParenthesis(list, text+"(", open+1, close, max)
	}
	if close < open {
		appendParenthesis(list, text+")", open, close+1, max)
	}
}
