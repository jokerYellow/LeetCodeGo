package _16_word_subsets

/*
https://leetcode.com/problems/word-subsets/

916. Word Subsets
Medium

160

50

Favorite

Share
We are given two arrays A and B of words.  Each word is a string of lowercase letters.

Now, say that word b is a subset of word a if every letter in b occurs in a, including multiplicity.  For example, "wrr" is a subset of "warrior", but is not a subset of "world".

Now say a word a from A is universal if for every b in B, b is a subset of a.

Return a list of all universal words in A.  You can return the words in any order.



Example 1:

Input: A = ["amazon","apple","facebook","google","leetcode"], B = ["e","o"]
Output: ["facebook","google","leetcode"]
Example 2:

Input: A = ["amazon","apple","facebook","google","leetcode"], B = ["l","e"]
Output: ["apple","google","leetcode"]
Example 3:

Input: A = ["amazon","apple","facebook","google","leetcode"], B = ["e","oo"]
Output: ["facebook","google"]
Example 4:

Input: A = ["amazon","apple","facebook","google","leetcode"], B = ["lo","eo"]
Output: ["google","leetcode"]
Example 5:

Input: A = ["amazon","apple","facebook","google","leetcode"], B = ["ec","oc","ceo"]
Output: ["facebook","leetcode"]


Note:

1 <= A.length, B.length <= 10000
1 <= A[i].length, B[i].length <= 10
A[i] and B[i] consist only of lowercase letters.
All words in A[i] are unique: there isn't i != j with A[i] == A[j].
*/

func wordSubsets(A []string, B []string) []string {
	lettersCount := make(map[rune]int, 26)
	for _, k := range B {
		for r, c := range maps(k) {
			if lettersCount[r] < c {
				lettersCount[r] = c
			}
		}
	}
	var rt []string
	for _, k := range A {
		amaps := maps(k)
		contain := true
		for r, c := range lettersCount {
			if amaps[r] < c {
				contain = false
				break
			}
		}
		if contain {
			rt = append(rt, k)
		}
	}
	return rt
}

func maps(word string) map[rune]int {
	lettersCount := make(map[rune]int)
	for _, r := range word {
		lettersCount[r] += 1
	}
	return lettersCount
}
