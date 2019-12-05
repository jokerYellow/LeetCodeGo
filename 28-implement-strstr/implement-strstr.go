package leetcode

func strStr(haystack string, needle string) int {
	if len(needle) == 0 {
		return 0
	}
	length := len(needle)
	heylength := len(haystack)
	for index, _ := range haystack {
		if index+length <= heylength && haystack[index:index+length] == needle {
			return index
		}
	}
	return -1
}
