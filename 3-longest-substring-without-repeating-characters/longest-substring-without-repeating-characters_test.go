package leetcode

import (
	"fmt"
	"testing"
)

func Test1(t *testing.T) {
	var s string
	var expect int
	var output int

	s = "pwwkew"
	expect = 3
	output = lengthOfLongestSubstring(s)
	fmt.Printf("expect:%d output:%d\n",expect,output)
	if output != expect{
		t.Fail()
	}
}


func Test2(t *testing.T) {
	var s string
	var expect int
	var output int

	s = "aaaaa"
	expect = 1
	output = lengthOfLongestSubstring(s)
	fmt.Printf("expect:%d output:%d\n",expect,output)
	if output != expect{
		t.Fail()
	}
}


func Test3(t *testing.T) {
	var s string
	var expect int
	var output int

	s = "ab"
	expect = 2
	output = lengthOfLongestSubstring(s)
	fmt.Printf("expect:%d output:%d\n",expect,output)
	if output != expect{
		t.Fail()
	}
}


func Test4(t *testing.T) {
	var s string
	var expect int
	var output int

	s = "dvdf"
	expect = 3
	output = lengthOfLongestSubstring(s)
	fmt.Printf("expect:%d output:%d\n",expect,output)
	if output != expect{
		t.Fail()
	}
}