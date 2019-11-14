package printBraces

import "fmt"

func printBraces(count int) []string {
	if count == 0 {
		return []string{}
	} else if count == 1 {
		return []string{"{", "}"}
	}
	var result []string
	items := printBracesI(count)
	for _, item := range items {
		if isValid(item) {
			result = append(result, item)
		}
	}
	return result
}

func printBracesI(count int) []string {
	if count == 0 {
		return []string{}
	} else if count == 1 {
		return []string{"{", "}"}
	}
	var items []string
	for _, item := range printBracesI(count - 1) {
		if isValidI(item) {
			items = append(items, fmt.Sprintf("%s}", item))
			items = append(items, fmt.Sprintf("%s{", item))
		}
	}
	return items
}

func isValidI(str string) bool {
	left := 0
	right := 0
	for _, b := range str {
		if b == '{' {
			left++
		} else {
			right++
		}
	}
	return left >= right
}

func isValid(str string) bool {
	var arr []rune
	for _, b := range str {
		if b == '}' {
			if len(arr) >= 1 && arr[len(arr)-1] == '{' {
				arr = arr[:len(arr)-1]
			} else {
				return false
			}
		} else {
			arr = append(arr, b)
		}
	}
	return len(arr) == 0
}

func printParentesis(n int) {
	if n > 0 {
		result := make([]byte, n)
		_printParenthesis(result, 0, n/2, 0, 0)
	}
}

func _printParenthesis(strs []byte, pos, n, open, close int) {
	if close == n {
		fmt.Println(string(strs))
	} else {
		if open > close {
			strs[pos] = '}'
			_printParenthesis(strs, pos+1, n, open, close+1)
		}
		if open < n {
			strs[pos] = '{'
			_printParenthesis(strs, pos+1, n, open+1, close)
		}
	}
}
