package leetcode

import (
	"math/rand"
)

/*
https://leetcode.com/problems/shuffle-an-array/

384. Shuffle an Array
Medium

246

600

Favorite

Share
Shuffle a set of numbers without duplicates.

Example:

// Init an array with set 1, 2, and 3.
int[] nums = {1,2,3};
Solution solution = new Solution(nums);

// Shuffle the array [1,2,3] and return its result. Any permutation of [1,2,3] must equally likely to be returned.
solution.shuffle();

// Resets the array back to its original configuration [1,2,3].
solution.reset();

// Returns the random shuffling of array [1,2,3].
solution.shuffle();
 */

type Solution struct {
	origin []int
}

func Constructor(nums []int) Solution {
	return Solution{nums}
}

/** Resets the array to its original configuration and return it. */
func (this *Solution) Reset() []int {
	return this.origin
}

/** Returns a random shuffling of the array. */
func (this *Solution) Shuffle() []int {

	clone := make([]int, len(this.origin))
	for i,v := range this.origin{
		clone[i] = v
	}

	length := len(clone)
	for i := range clone {
		d := rand.Intn(length)
		t := clone[d]
		clone[d] = clone[i]
		clone[i] = t
	}
	return clone
}


/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.Reset();
 * param_2 := obj.Shuffle();
 */
