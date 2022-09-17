package dp

import "math"

// Recursion.

func ClimbStairsRec(nStairs, inc int) int {
	mod := 1000000007

	if inc == nStairs {
		return 1
	}

	if inc > nStairs {
		return 0
	}

	return (ClimbStairsRec(nStairs, inc+1) + ClimbStairsRec(nStairs, inc+2)) % mod
}

//  Original Question.   Need!
func MinCostClimbingStairs(cost []int) int {
	n := len(cost)

	ans := math.Min(float64(solve(cost, n-1)), float64(solve(cost, n-2)))
	return int(ans)
}
func solve(cost []int, n int) int {
	// base case
	if n == 0 {
		return cost[0]
	}
	if n == 1 {
		return cost[1]
	}

	ans := cost[n] + int(math.Min(float64(solve(cost, n-1)), float64(solve(cost, n-2))))
	return ans
}
