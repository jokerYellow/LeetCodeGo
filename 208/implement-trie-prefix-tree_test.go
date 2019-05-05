package leetcode

import "testing"

func Test1(t *testing.T) {
	 obj := Constructor()
	 obj.Insert("a")
	 param_2 := obj.Search("a")
	 param_3 := obj.StartsWith("a")
	 if param_2 != true || param_3 != true{
	 	t.Fail()
	 }

}