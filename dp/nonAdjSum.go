package dp

import (
	"math"
)

// Max sum of non-adjcent element.  || House robber - 1

// recursion based solution.
func MaxSumNonAdjMain(arr []int) int {
	// n := len(arr)
	//TYPE 1 recursion based solution.
	// ans := recSolve(arr, n-1)

	// TYPE 2 RECURSION WITH MEMORISATION.
	// step 1 => create dp arr.
	// dp := []int{}
	// for i := 0; i < n; i++ {
	// 	dp = append(dp, -1)
	// }

	// TYPE 3 TABULATION METHOD.

	return tabSolve(arr)
}

func tabSolve(arr []int) int {
	n := len(arr)
	// Create dp array
	dp := make([]int, n)

	// base case
	dp[0] = arr[0]
	for i := 1; i < n; i++ {
		inc := arr[i]
		if i > 1 { // Apply condition
			inc += dp[i-2]
		}
		excl := dp[i-1] + 0

		dp[i] = int(math.Max(float64(inc), float64(excl)))
	}

	return dp[n-1]
}
func memoSolve(arr, dp []int, n int) int {
	// base case
	if n < 0 {
		return 0
	}

	if n == 0 {
		return arr[0]
	}

	// STEP 3. check in dp array call are there or not

	if dp[n] != -1 {
		return dp[n]
	}

	// include => jump 2 steps and include the current number.
	inc := memoSolve(arr, dp, n-2) + arr[n]
	// exclude => jump 1 step but not inclue the current number.
	exc := memoSolve(arr, dp, n-1) + 0

	// STEP 2. ASSIGN DP ARRAY
	dp[n] = int(math.Max(float64(inc), float64(exc)))

	return dp[n]
}
func recSolve(arr []int, n int) int {
	// base case
	if n < 0 {
		return 0
	}

	if n == 0 {
		return arr[0]
	}

	// include => jump 2 steps and include the current number.
	inc := recSolve(arr, n-2) + arr[n]
	// exclude => jump 1 step but not inclue the current number.
	exc := recSolve(arr, n-1) + 0

	return int(math.Max(float64(inc), float64(exc)))

}
