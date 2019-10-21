package quicksort

func quisksort(arr []int) []int {
	if len(arr) == 0 {
		return arr
	}
	out := arr
	part(out, 0, 0, len(arr)-1)
	return out
}

//increase
//move the arr,make left of the index is littler ,right is bigger
func part(arr []int, index, leftBound, rightBound int) {
	if leftBound >= rightBound {
		return
	}
	left := leftBound
	right := rightBound
	for left < right {
		for left < index {
			if arr[left] < arr[index] {
				left++
			} else {
				arr[left], arr[index] = arr[index], arr[left]
				index = left
				left++
				break
			}
		}
		for right > index {
			if arr[right] >= arr[index] {
				right--
			} else {
				arr[right], arr[index] = arr[index], arr[right]
				index = right
				right--
				break
			}
		}
	}
	part(arr, leftBound, leftBound, index-1)
	part(arr, index+1, index+1, rightBound)
}
