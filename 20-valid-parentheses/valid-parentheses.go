package _0_valid_parentheses

/*
20. Valid Parentheses
Easy

2922

145

Favorite

Share
Given a string containing just the characters '(', ')', '{', '}', '[' and ']', determine if the input string is valid.

An input string is valid if:

Open brackets must be closed by the same type of brackets.
Open brackets must be closed in the correct order.
Note that an empty string is also considered valid.

Example 1:

Input: "()"
Output: true
Example 2:

Input: "()[]{}"
Output: true
Example 3:

Input: "(]"
Output: false
Example 4:

Input: "([)]"
Output: false
Example 5:

Input: "{[]}"
Output: true
*/
type stack struct {
	arr []interface{}
}

func (s *stack) pop() interface{} {
	t := s.top()
	if t != nil {
		s.arr = s.arr[:len(s.arr)-1]
	}
	return t
}

func (s *stack) top() interface{} {
	if len(s.arr) > 0 {
		top := s.arr[len(s.arr)-1]
		return top
	}
	return nil
}

func (s *stack) push(item interface{}) {
	s.arr = append(s.arr, item)
}

//'(', ')', '{', '}', '[' and ']',
func isInPair(l, r rune) bool {
	if l == '(' && r == ')' {
		return true
	} else if l == '{' && r == '}' {
		return true
	} else if l == '[' && r == ']' {
		return true
	}
	return false
}

func isValid(s string) bool {
	st := stack{}
	for _, v := range s {
		if st.top() != nil && isInPair(st.top().(rune), v) {
			st.pop()
		} else {
			st.push(v)
		}
	}
	return st.top() == nil
}
