package leetcode

import (
	"testing"
	"github.com/jokerYellow/leetcode/utils"
)

type testCase struct {
	strs []string 
	expect [][]string
}

func Test(t*testing.T){
	cases := []testCase{
		{[]string{"eat","tea","tan","ate","nat","bat"},[][]string{{"eat","tea","ate"},{"tan","nat"},{"bat"}}},
		{[]string{"吃","吃"},[][]string{{"吃","吃"},}},
	}
	for _,item := range cases{
		output := groupAnagrams(item.strs) 
		if !utils.CheckStringsArrEqual(output,item.expect){
			t.Fail()
			utils.Print(item.expect,output)
		}
	}
}