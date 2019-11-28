package leetcode

/*
1227. Airplane Seat Assignment Probability
Medium

66

92

Favorite

Share
n passengers board an airplane with exactly n seats. The first passenger has lost the ticket and picks a seat randomly. But after that, the rest of passengers will:

Take their own seat if it is still available,
Pick other seats randomly when they find their seat occupied
What is the probability that the n-th person can get his own seat?



Example 1:

Input: n = 1
Output: 1.00000
Explanation: The first person can only get the first seat.
Example 2:

Input: n = 2
Output: 0.50000
Explanation: The second person has a probability of 0.5 to get the second seat (when first person gets the first seat).


Constraints:

1 <= n <= 10^5
 */


//f(2) = 1/2
//f(3) = 1/3 + (3-2)/3*f(2) = 1/2

//f(n+1) = 1/(n+1) + (n+1-2)/(n+1) * f(n)
//suppose f(n) = 1/2, f(n+1) = (1 + (n+1-2)*0.5)/(n+1) = 1/2
//so when n > 1, f(n) = 1/1
func nthPersonGetsNthSeat(n int) float64 {
	if n == 1 {
		return 1
	}
	return 0.5
}

func nthPersonGetsNthSeat1(n int) float64 {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}else if n == 2 {
		return 0.5
	}
	return 1/float64(n) + (float64(n) -2)/float64(n) * (nthPersonGetsNthSeat(n-1))
}