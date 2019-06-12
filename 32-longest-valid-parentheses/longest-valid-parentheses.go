package _2_longest_valid_parentheses

/*
https://leetcode.com/problems/longest-valid-parentheses/
32. Longest Valid Parentheses
Hard

1865

90

Favorite

Share
Given a string containing just the characters '(' and ')', find the length of the longest valid (well-formed) parentheses substring.

Example 1:

Input: "(()"
Output: 2
Explanation: The longest valid parentheses substring is "()"
Example 2:

Input: ")()())"
Output: 4
Explanation: The longest valid parentheses substring is "()()"
*/

type item struct {
	index int
	value rune
}

type stack struct {
	arr []item
}

func (s *stack) pop() *item {
	t := s.top()
	if t != nil {
		s.arr = s.arr[:len(s.arr)-1]
	}
	return t
}

func (s *stack) top() *item {
	if len(s.arr) > 0 {
		top := s.arr[len(s.arr)-1]
		return &top
	}
	return nil
}

func (s *stack) isEmpty() bool {
	return len(s.arr) == 0
}

func (s *stack) push(i item) {
	s.arr = append(s.arr, i)
}

func isPair(l, r rune) bool {
	return l == '(' && r == ')'
}

func longestValidParentheses(s string) int {
	longest := 0
	st := stack{}
	count := 0
	longs := make([]int, len(s))
	for index, v := range s {
		if st.top() != nil && isPair(st.top().value, v) {
			count += 2
			lessLeftIndex := st.top().index - 1
			if lessLeftIndex >= 0 {
				count += longs[lessLeftIndex]
			}
			lessRightIndex := index - 1
			if lessRightIndex >= 0 {
				count += longs[lessRightIndex]
			}
			longs[index] = count
			if longest < count {
				longest = count
			}
			count = 0
			st.pop()
		} else {
			st.push(item{index, v})
		}
	}
	return longest
}
