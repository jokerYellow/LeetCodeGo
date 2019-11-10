package searchMaxN

func searchMaxN(arr []int, index int) int {
	if len(arr) <= index {
		return -1
	}
	left := 0
	right := len(arr) - 1
	searchIndex := 0
	index--
	for {
		result := part(arr, left, left, right)
		if result < index {
			left = result + 1
		} else if result == index {
			searchIndex = result
			break
		} else if result > index {
			right = result - 1
		}
	}
	return arr[searchIndex]
}

func part(arr []int, index, leftBound, rightBound int) int {
	left := leftBound
	right := rightBound
	for left < right {
		for left < index {
			if arr[left] >= arr[index] {
				left++
			} else {
				arr[left], arr[index] = arr[index], arr[left]
				index = left
				break
			}
		}
		for right > index {
			if arr[right] <= arr[index] {
				right--
			} else {
				arr[right], arr[index] = arr[index], arr[right]
				index = right
				break
			}
		}
	}
	return index
}
