package powx_n

import "fmt"

func myPow(x float64, n int) float64 {
	if x == 0 || x == 1 {
		return x
	}
	if n < 0 {
		return 1 / myPow(x, -n)
	}
	var result float64 = 1
	i := 1
	for {
		fmt.Println(result, i)
		newResult := result * result
		if 2*i > n {
			result = result * myPow(x, n-i)
			break
		} else if 2*i == n {
			return newResult
		} else {
			result = newResult
			i = 2 * i
		}
	}
	return result
}
