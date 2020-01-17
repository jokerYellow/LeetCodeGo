package isPowerOfFour

func isPowerOfFour(num int) bool {
     return isPowerOf(num,4)
}

func isPowerOf(n int, v int) bool {
	if n <= 0 {
		return false
	}
	for n != 1{
		if (n/v)*v != n {
			return false
		}
		n = n / v
	}
	return true
}