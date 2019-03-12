package _003

/*
1003. Check If Word Is Valid After Substitutions
Medium

31

49

Favorite

Share
We are given that the string "abc" is valid.

From any valid string V, we may split V into two pieces X and Y such that X + Y (X concatenated with Y) is equal to V.  (X or Y may be empty.)  Then, X + "abc" + Y is also valid.

If for example S = "abc", then examples of valid strings are: "abc", "aabcbc", "abcabc", "abcabcababcc".  Examples of invalid strings are: "abccba", "ab", "cababc", "bac".

Return true if and only if the given string S is valid.



Example 1:

Input: "aabcbc"
Output: true
Explanation:
We start with the valid string "abc".
Then we can insert another "abc" between "a" and "bc", resulting in "a" + "abc" + "bc" which is "aabcbc".
Example 2:

Input: "abcabcababcc"
Output: true
Explanation:
"abcabcabc" is valid after consecutive insertings of "abc".
Then we can insert "abc" before the last letter, resulting in "abcabcab" + "abc" + "c" which is "abcabcababcc".
Example 3:

Input: "abccba"
Output: false
Example 4:

Input: "cababc"
Output: false


Note:

1 <= S.length <= 20000
S[i] is 'a', 'b', or 'c'

*/

type stack struct {
	items []rune
}

func (s *stack) pop() bool {
	length := len(s.items)
	if length >= 3 {
		t := s.items[length-3 : length]
		if t[0] == 'a' && t[1] == 'b' && t[2] == 'c' {
			s.items = s.items[:length-3]
			return true
		}
	}
	return false
}

func (s *stack) isClean() bool {
	return len(s.items) == 0
}

func (s *stack) push(b rune) {
	s.items = append(s.items, b)
}

func isValid(S string) bool {
	s := new(stack)
	for _, v := range S {
		s.push(v)
		if v == 'c' && s.pop() == false {
			return false
		}
	}
	return s.isClean()
}
