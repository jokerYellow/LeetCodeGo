package leetcode
/*
https://leetcode.com/problems/group-anagrams/
49. Group Anagrams
Medium

2325

138

Add to List

Share
Given an array of strings, group anagrams together.

Example:

Input: ["eat", "tea", "tan", "ate", "nat", "bat"],
Output:
[
  ["ate","eat","tea"],
  ["nat","tan"],
  ["bat"]
]
Note:

All inputs will be in lowercase.
The order of your output does not matter.
*/
import "sort"

type sortBytes []byte 

func (s sortBytes) Less(i,j int)bool{
	return s[i] < s[j]
}

func (s sortBytes)Swap(i,j int){
	s[i],s[j] = s[j],s[i]
}

func (s sortBytes) Len() int{
	return len(s)
} 

func groupAnagrams(strs []string)[][]string{
	m := map[string][]string{}
	for _,item := range strs{
		b := []byte(item)
		sort.Sort(sortBytes(b))
		s := string(b)
		if v,ok := m[s];ok{
			m[s] = append(v,item)
		}else{
			m[s] = []string{item}
		}
	}
	var values [][]string
	for _,v := range m {
		values = append(values,v)
	}
	return values
}