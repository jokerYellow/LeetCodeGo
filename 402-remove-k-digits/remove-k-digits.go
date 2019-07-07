package _02_remove_k_digits

import (
	"strings"
)

/*
https://leetcode.com/problems/remove-k-digits/

402. Remove K Digits
Medium

908

59

Favorite

Share
Given a non-negative integer num represented as a string, remove k digits from the number so that the new number is the smallest possible.

Note:
The length of num is less than 10002 and will be ≥ k.
The given num does not contain any leading zero.
Example 1:

Input: num = "1432219", k = 3
Output: "1219"
Explanation: Remove the three digits 4, 3, and 2 to form the new number 1219 which is the smallest.
Example 2:

Input: num = "10200", k = 1
Output: "200"
Explanation: Remove the leading 1 and the number is 200. Note that the output must not contain leading zeroes.
Example 3:

Input: num = "10", k = 2
Output: "0"
Explanation: Remove all the digits from the number and it is left with nothing which is 0.
*/

type stack struct {
	items []rune
}

func newStack() *stack {
	s := &stack{}
	s.items = []rune{}
	return s
}

func (s *stack) pop() rune {
	top := s.top()
	s.items = s.items[:s.count()-1]
	return top
}

func (s *stack) count() int {
	return len(s.items)
}

func (s *stack) isEmpty() bool {
	return s.count() == 0
}

func (s *stack) top() rune {
	if len(s.items) == 0 {
		return 0
	}
	return s.items[s.count()-1]
}

func (s *stack) push(item rune) {
	s.items = append(s.items, item)
}

func removeKdigits(num string, k int) string {
	if len(num) <= k {
		return "0"
	}
	if len(num) == 0 {
		return "0"
	}
	s := newStack()
	count := k
	for _, k := range num {
		v := k
		if s.isEmpty() {
			s.push(v)
			continue
		}

		for s.top() > v && count > 0 {
			s.pop()
			count -= 1
		}

		s.push(v)
	}
	length := len(num) - k
	var zero rune = 48
	b := strings.Builder{}

	for index, v := range s.items {
		if index >= length {
			break
		}
		if v == zero && b.Len() == 0 {
			continue
		}
		b.WriteRune(v)
	}
	rt := b.String()
	if rt == "" {
		rt = "0"
	}
	return rt
}
