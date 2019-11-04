package leetcode

type heap struct {
	items []int
}

func left(index int) int {
	return index * 2
}

func right(index int) int {
	return left(index) + 1
}

func newHeap(items []int) *heap {
	h := new(heap)
	h.items = items
	h.buildMaxHeap()
	return h
}

func (h *heap) buildMaxHeap() {
	length := len(h.items)
	for i := length / 2; i >= 0; i-- {
		h.maxHeap(i)
	}
}

func (h *heap) append(item int) {
	h.items = append(h.items, item)
}

func (h *heap) maxHeap(index int) {
	length := len(h.items)
	leftIndex := left(index)
	rightIndex := right(index)
	maxIndex := index
	if leftIndex < length && h.items[leftIndex] > h.items[maxIndex] {
		maxIndex = leftIndex
	}
	if rightIndex < length && h.items[rightIndex] > h.items[maxIndex] {
		maxIndex = rightIndex
	}
	if maxIndex != index {
		h.items[index], h.items[maxIndex] = h.items[maxIndex], h.items[index]
		h.maxHeap(maxIndex)
	}
}

func (h *heap) removeTop() int {
	if len(h.items) == 0 {
		return -1
	}
	top := h.items[0]
	h.items[0] = h.items[len(h.items)-1]
	h.items = h.items[:len(h.items)-1]
	h.maxHeap(0)
	return top
}

func findKthLargest(nums []int, k int) int {
	h := newHeap(nums)
	count := k
	last := 0
	for count > 0 {
		c := h.removeTop()
		last = c
		count--
	}
	return last
}
