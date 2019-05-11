package __reverse_integer

func reverse(x int) int {
	if isOverFlow(x) {
		return 0
	}
	var rt int
	for x != 0 {
		t := x % 10
		rt = rt*10 + t
		x = x / 10
	}
	if isOverFlow(rt) {
		return 0
	}
	return rt
}

func isOverFlow(x int) bool {
	upperBounds := 2147483647  //2**31 -1
	lowerBounds := -2147483648 //-2*31
	if x > upperBounds || x < lowerBounds {
		return true
	}
	return false
}
