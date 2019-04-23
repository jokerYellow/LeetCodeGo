package leetcode

import "fmt"

/*
https://leetcode.com/problems/4sum-ii/
454. 4Sum II
Medium

637

56

Favorite

Share
Given four lists A, B, C, D of integer values, compute how many tuples (i, j, k, l) there are such that A[i] + B[j] + C[k] + D[l] is zero.

To make problem a bit easier, all A, B, C, D have same length of N where 0 ≤ N ≤ 500. All integers are in the range of -228 to 228 - 1 and the result is guaranteed to be at most 231 - 1.

Example:

Input:
A = [ 1, 2]
B = [-2,-1]
C = [-1, 2]
D = [ 0, 2]

Output:
2

Explanation:
The two tuples are:
1. (0, 0, 0, 1) -> A[0] + B[0] + C[0] + D[1] = 1 + (-2) + (-1) + 2 = 0
2. (1, 1, 0, 0) -> A[1] + B[1] + C[0] + D[0] = 2 + (-1) + (-1) + 0 = 0
 */

func fourSumCount(A []int, B []int, C []int, D []int) int {
	fmt.Println(len(A))
	Dmap := map[int]int{}
	for _, v := range D {
		Dmap[v] ++
	}

	count := 0
	for _, a := range A {
		for _, b := range B {
			for _, c := range C {
				t := 0 - (a + b + c)
				if Dmap[t] != 0 {
					count += Dmap[t]
				}
			}
		}
	}

	return count
}
