package powerOfThree

func isPowerOfThree(n int) bool {
	return isPowerOf(n, 3)
}

func isPowerOf(n int, v int) bool {
	if n <= 0 {
		return false
	}
	for n != 1 {
		if (n/v)*v != n {
			return false
		}
		n = n / v
	}
	return true
}
