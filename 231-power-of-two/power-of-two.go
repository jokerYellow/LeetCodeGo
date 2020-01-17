package poweroftwo

func isPowerOfTwo(n int) bool {
	return isPowerOf(n, 1)
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

func _isPowerOfTwo(n int) bool {
	v := 1
	for i := 0; i < 32; i++ {
		if n == v {
			return true
		}
		v = v * 2
	}
	return false
}
