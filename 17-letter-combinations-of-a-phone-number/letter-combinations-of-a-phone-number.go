package leetcode

func letterCombinations(digits string) []string {
	runes := _letterCombinations([]rune(digits))
	rt := []string{}
	for _, r := range runes {
		rt = append(rt, string(r))
	}
	return rt
}

func _letterCombinations(digits []rune) [][]rune {
	letters := map[rune][]rune{
		'2': []rune{'a', 'b', 'c'},
		'3': []rune{'d', 'e', 'f'},
		'4': []rune{'g', 'h', 'i'},
		'5': []rune{'j', 'k', 'l'},
		'6': []rune{'m', 'n', 'o'},
		'7': []rune{'p', 'q', 'r', 's'},
		'8': []rune{'t', 'u', 'v'},
		'9': []rune{'w', 'x', 'y', 'z'},
	}
	if len(digits) == 0 {
		return [][]rune{}
	}
	if runes, ok := letters[digits[0]]; ok {
		result := [][]rune{}
		last := _letterCombinations(digits[1:])
		for _, r := range runes {
			if len(last) == 0 {
				result = append(result, []rune{r})
			} else {
				for _, l := range last {
					t := []rune{r}
					t = append(t, l...)
					result = append(result, t)
				}
			}
		}
		return result
	}
	return _letterCombinations(digits[1:])
}
