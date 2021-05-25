package __palindrome_number

func isPalindrome(x int) bool {
	if x < 0 || (x > 0 && x%10 == 0) {
		return false
	}
	var reverse int
	for x > reverse {
		reverse = 10*reverse + x%10
		x = x / 10
	}
	return reverse == x || reverse/10 == x
}

func isPalindromeNormal(x int) bool {
	if x < 0 {
		return false
	}
	t := x
	var rt int
	for t != 0 {
		rt = 10*rt + t%10
		t = t / 10
	}
	return rt == x
}
