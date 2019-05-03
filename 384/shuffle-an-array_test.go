package leetcode

import (
	"fmt"
	"testing"
)

func Test1(t *testing.T) {
	obj1 := Constructor([]int{1, 2, 3})
	fmt.Println(obj1.Reset())

	for i := 0; i < 10; i++ {
		fmt.Println(obj1.Shuffle())
	}

}
