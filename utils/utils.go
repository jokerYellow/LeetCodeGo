package utils

import "fmt"

func EqualStrings(l, r []string) bool {
	if len(l) != len(r) {
		return false
	}
	i := 0
	for i < len(l) {
		if l[i] != r[i] {
			return false
		}
		i++
	}
	return true
}

func CheckEqual(a, b [][]int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if CheckEqualArr(v, b[i]) == false {
			return false
		}
	}
	return true
}

func CheckElementsEqual(nums1, nums2 [][]int) bool {
	if len(nums1) != len(nums2) {
		return false
	}
	for _, v := range nums1 {
		if !checkElementsContain(nums2, v) {
			return false
		}
	}
	return true
}

func checkElementsContain(nums1 [][]int, element []int) bool {
	flag := false
	for _, item := range nums1 {
		if CheckEqualArr(item, element) {
			flag = true
			break
		}
	}
	return flag
}

//item in a ,is alse in b
func CheckItemsEqual(a, b [][]int) bool {
	if len(a) != len(b) {
		return false
	}
	for _, ia := range a {
		contain := false
		for _, ib := range b {
			if CheckEqualArr(ia, ib) {
				contain = true
				break
			}
		}
		if !contain {
			return false
		}
	}
	return true
}

func CheckEqualArr(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if b[i] != v {
			return false
		}
	}
	return true
}

func CheckStringsEqualArr(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if b[i] != v {
			return false
		}
	}
	return true
}

func Print(expect, output interface{}) {
	fmt.Printf("expect:%s\noutput:%s\n", transform(expect), transform(output))
}

func transform(value interface{}) string {
	switch v := value.(type) {
	case []byte:
		return fmt.Sprintf("%s:%v", v, v)
	case []rune:
		return fmt.Sprintf("%s:%v", string(v), v)
	case *ListNode:
		return v.description()
	default:
		return fmt.Sprintf("%v", v)
	}
}

func Contain(arr []string, s string) bool {
	for _, v := range arr {
		if v == s {
			return true
		}
	}
	return false
}

func CheckStringsArrEqual(str1, str2 [][]string) bool {
	if len(str1) != len(str2) {
		return false
	}
	for i := 0; i < len(str1); i++ {
		if !CheckStringsEqual(str1[i], str2[i]) {
			return false
		}
	}
	return true
}

func CheckStringsEqual(str1, str2 []string) bool {
	if len(str1) != len(str2) {
		return false
	}
	for i := 0; i < len(str1); i++ {
		if str1[i] != str2[i] {
			return false
		}
	}
	return true
}
