package quicksort

func quisksort(arr []int) []int {
	if len(arr) == 0 {
		return arr
	}
	quickSort(arr, 0, len(arr)-1)
	return arr
}

func quickSort(arr []int, start, end int) {
	if start >= end {
		return
	}
	index := part(arr, start, start, len(arr)-1)
	if index-1 > start {
		quickSort(arr, start, index-1)
	}
	if index+1 < end {
		quickSort(arr, index+1, end)
	}
}

//increase
//move the arr,make left of the index is littler ,right is bigger
func part(arr []int, index, leftBound, rightBound int) int {
	left := leftBound
	right := rightBound
	for left < right {
		if left < index {
			if arr[left] <= arr[index] {
				left++
			} else {
				arr[left], arr[index] = arr[index], arr[left]
				index = left
			}
		}
		if index < right {
			if arr[right] >= arr[index] {
				right--
			} else {
				arr[right], arr[index] = arr[index], arr[right]
				index = right
			}
		}
	}
	return index
}
