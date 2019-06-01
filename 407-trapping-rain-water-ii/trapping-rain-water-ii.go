package _07_trapping_rain_water_ii

import "errors"

/*
https://leetcode.com/problems/trapping-rain-water-ii/
407. Trapping Rain Water II
Hard

672

17

Favorite

Share
Given an m x n matrix of positive integers representing the height of each unit cell in a 2D elevation map, compute the volume of water it is able to trap after raining.



Note:

Both m and n are less than 110. The height of each unit cell is greater than 0 and is less than 20,000.



Example:

Given the following 3x6 height map:
[
  [1,4,3,1,3,2],
  [3,2,1,3,2,4],
  [2,3,3,2,3,1]
]

Return 4.


The above image represents the elevation map [[1,4,3,1,3,2],[3,2,1,3,2,4],[2,3,3,2,3,1]] before the rain.





After the rain, water is trapped between the blocks. The total volume of water trapped is 4.


*/
/*
heap
1 2 3 4 3 2 1 2 2]

        1
    2      3
  4   3  2   1
 2 2
		4
	  3    3
     2 2  2 1
    1 2
*/

type heap struct {
	heapSize int
	items    []int
}

func left(i int) int {
	return 2 * i
}

func right(i int) int {
	return 2*i + 1
}

func (this *heap) index(index int) (int, error) {
	if len(this.items) < index {
		return 0, errors.New("except range")
	} else {
		return this.items[index-1], nil
	}
}

func (this *heap) exchange(index1, index2 int) {
	index1 -=1
	index2 -=1
	this.items[index1], this.items[index2] = this.items[index2], this.items[index1]
}

func (this *heap) maxHeapify(i int) {
	l := left(i)
	r := right(i)
	if l < len(this.items){
		this.maxHeapify(l)
	}
	if r < len(this.items){
		this.maxHeapify(r)
	}
	currentValue, err := this.index(i)
	if err != nil {
		return
	}
	leftValue, err := this.index(l)
	if err == nil {
		if leftValue > currentValue {
			this.exchange(l, i)
		}
	}
	currentValue, _ = this.index(i)
	rightValue, err := this.index(r)
	if err == nil {
		if rightValue > currentValue {
			this.exchange(r, i)
		}
	}
	this.maxHeapify(l)
	this.maxHeapify(r)
}

func (this *heap) buildMaxHeap() {

}

func (this *heap) maxHeapInsert() {

}

func (this *heap) heapExtractMax() {

}

func (this *heap) heapIncreaseKey() {

}

func (this *heap) heapMaximum() {

}

type cell struct {
	row    int
	column int
	height int
}

func trapRainWater(heightMap [][]int) int {
	if len(heightMap) == 0 {
		return 0
	}
	for i := 0; i < len(heightMap); i++ {

	}
	return 0
}
