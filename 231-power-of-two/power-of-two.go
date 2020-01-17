package poweroftwo

func isPowerOfTwo(n int) bool {
	if n <= 0 {
		return false
	}
	for n != 1 {
		if n&1 == 1 {
			return false
		}
		n = n >> 1
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
