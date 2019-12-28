package powx_n

func myPow(x float64, n int) float64 {
	if n == 0 {
		return 1
	}
	if x == 0 || x == 1 || n == 1 {
		return x
	}
	if n < 0 {
		return 1 / myPow(x, -n)
	}
	var result float64 = x
	i := 1
	for 2*i <= n {
		result = result * result
		i = 2 * i
	}
	return result * myPow(x, n-i)
}
