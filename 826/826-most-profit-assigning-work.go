package _26

/*
826. Most Profit Assigning Work
Medium

191

29

Favorite

Share
We have jobs: difficulty[i] is the difficulty of the ith job, and profit[i] is the profit of the ith job.

Now we have some workers. worker[i] is the ability of the ith worker, which means that this worker can only complete a job with difficulty at most worker[i].

Every worker can be assigned at most one job, but one job can be completed multiple times.

For example, if 3 people attempt the same job that pays $1, then the total profit will be $3.  If a worker cannot complete any job, his profit is $0.

What is the most profit we can make?

Example 1:

Input: difficulty = [2,4,6,8,10], profit = [10,20,30,40,50], worker = [4,5,6,7]
Output: 100
Explanation: Workers are assigned jobs of difficulty [4,4,6,6] and they get profit of [20,20,30,30] seperately.
Notes:

1 <= difficulty.length = profit.length <= 10000
1 <= worker.length <= 10000
difficulty[i], profit[i], worker[i]  are in range [1, 10^5]

*/
func maxProfitAssignment(difficulty []int, profit []int, worker []int) int {

	findProfit := func(w int) int {
		maxIndex := -1
		maxProfit := -1
		for i, v := range difficulty {
			if v <= w {
				if maxProfit < profit[i] {
					maxIndex = i
					maxProfit = profit[i]
				}
			}
		}
		if maxIndex <= (len(profit)-1) && maxIndex >= 0 {
			return profit[maxIndex]
		}
		return 0
	}
	rt := 0
	for _, work := range worker {
		rt += findProfit(work)
	}

	return rt
}
