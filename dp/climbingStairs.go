package dp

import (
	"math"
)

// Recursion.

// func ClimbStairsRec(nStairs, inc int) int {
// 	mod := 1000000007

// 	if inc == nStairs {
// 		return 1
// 	}

// 	if inc > nStairs {
// 		return 0
// 	}

// 	return (ClimbStairsRec(nStairs, inc+1) + ClimbStairsRec(nStairs, inc+2)) % mod
// }

// Original Question.   Need!
func MinCostClimbingStairs(cost []int) int {
	n := len(cost)

	// ans := math.Min(float64(solve(cost, n-1)), float64(solve(cost, n-2)))
	// return int(ans)

	// tabulation

	dp := make([]int, n+1)

	ans := tabMinCostClimbingSolve(cost, n, dp)

	return ans

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

func tabMinCostClimbingSolve(cost []int, n int, dp []int) int {
	if n == 0 {
		return cost[0]
	}
	if n == 1 {
		return cost[1]
	}

	if dp[n] != 0 {
		return dp[n]
	}

	dp[n] = cost[n] + int(math.Min(float64(tabMinCostClimbingSolve(cost, n-1, dp)), float64(tabMinCostClimbingSolve(cost, n-2, dp))))

	return dp[n]
}

func bottomUpCostClimbing(cost []int, n int) int {
	dp := make([]int, n+1)

	dp[0] = cost[0]
	dp[1] = cost[1]

	for i := 2; i < n; i++ {
		dp[i] = cost[i] + int(math.Min(float64(dp[i-1]), float64(dp[i-2])))
	}

	return int(math.Min(float64(dp[n-1]), float64(dp[n-2])))
}
