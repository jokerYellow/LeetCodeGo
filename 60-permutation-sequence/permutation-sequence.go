package leetcode

/*
https://leetcode.com/problems/permutation-sequence/
60. Permutation Sequence
Medium

1105

279

Add to List

Share
The set [1,2,3,...,n] contains a total of n! unique permutations.

By listing and labeling all of the permutations in order, we get the following sequence for n = 3:

"123"
"132"
"213"
"231"
"312"
"321"
Given n and k, return the kth permutation sequence.

Note:

Given n will be between 1 and 9 inclusive.
Given k will be between 1 and n! inclusive.
Example 1:

Input: n = 3, k = 3
Output: "213"
Example 2:

Input: n = 4, k = 9
Output: "2314"

"123"
"132"
"213"
"231"
"312"
"321"

*/
func getPermutation(n int, k int) string {
	arr := make([]byte, n)
	for i := 0; i < n; i++ {
		arr[i] = '0' + byte(i+1)
	}
	return string(_getPermutation(arr, k))
}
//通过计算出首位char，来缩小问题规模。
func _getPermutation(arr []byte, k int) []byte {
	if k <= 1 {
		return arr
	}
	if k == 2 && len(arr) == 2 {
		return []byte{arr[1], arr[0]}
	}
	count := v(len(arr) - 1)
	removeIndex := k / count
	if k%count == 0 {
		removeIndex--
	}
	var buf []byte
	buf = append(buf, arr[removeIndex])
	var newArr []byte
	newArr = append(newArr, arr[:removeIndex]...)
	newArr = append(newArr, arr[removeIndex+1:]...)
	buf = append(buf, _getPermutation(newArr, k-count*removeIndex)...)
	return buf
}

func v(n int) int {
	if n == 1 {
		return 1
	}
	return n * v(n-1)
}
