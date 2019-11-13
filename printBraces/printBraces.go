package printBraces

import "fmt"

func printBraces(count int) []string {
	if count == 0 {
		return []string{}
	} else if count == 1 {
		return []string{"{", "}"}
	}
	var result []string
	for _, item := range printBracesI(count) {
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
		items = append(items, fmt.Sprintf("%s}", item))
		items = append(items, fmt.Sprintf("%s{", item))
	}
	return items
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
