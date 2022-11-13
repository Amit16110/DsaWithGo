package striing

import (
	"fmt"
	"math"
)

func solve(a, b string, i, j int) int {
	// base case
	if i == len(a) {
		return 0
	}
	if j == len(b) {
		return 0
	}

	ans := 0
	// check if the string a[i] == b[j] is true means add the 1 with recursion call.
	if a[i] == b[j] { //match condition
		return 1 + solve(a, b, i+1, j+1)
	} else {
		// here the if a[i] != b[j] so we have two options
		// we can increment the i pointer and compare with whole j or
		// we can increment the j pointer and compare with whole i. and store the max of both calls.
		ans = int(math.Max(float64(solve(a, b, i+1, j)), float64(solve(a, b, i, j+1))))
	}
	return ans
}

func memoSolve(a, b string, i, j int, dp [][]int) int {
	if i == len(a) {
		return 0
	}
	if j == len(b) {
		return 0
	}

	if dp[i][j] != 0 {
		return dp[i][j]
	}
	fmt.Println("dp =>", dp)
	ans := 0
	if a[i] == b[j] { //match condition
		return 1 + memoSolve(a, b, i+1, j+1, dp)
	} else {
		// here the if a[i] != b[j] so we have two options
		// we can increment the i pointer and compare with whole j or
		// we can increment the j pointer and compare with whole i. and store the max of both calls.
		ans = int(math.Max(float64(memoSolve(a, b, i+1, j, dp)), float64(memoSolve(a, b, i, j+1, dp))))
	}
	dp[i][j] = ans
	return dp[i][j]
}

func bottomUpSolve(a, b string) int {
	//tabulation method

	// Step 1. => first create the dp array in the different function.
	//Step 2. => Go to recursive solution and check the base case and analyis what you need
	//to do in the base.
	//Step 3. => change the function call with dp store.
	n := len(a)
	nn := len(b)

	dp := make([][]int, n+1)
	for i, _ := range dp {
		dp[i] = make([]int, nn+1)
	}

	// base case no need cause dp array already init with 0.
	for i := len(a) - 1; i >= 0; i-- {
		for j := len(b) - 1; j >= 0; j-- {

			ans := 0
			if a[i] == b[j] { //match condition
				ans = 1 + dp[i+1][j+1]
			} else {
				// here the if a[i] != b[j] so we have two options
				// we can increment the i pointer and compare with whole j or
				// we can increment the j pointer and compare with whole i. and store the max of both calls.
				ans = int(math.Max(float64(dp[i+1][j]), float64(dp[i][j+1])))
			}
			dp[i][j] = ans

		}
	}
	fmt.Println("bottom up", dp)
	return dp[0][0]
}

func spaceOptimization(a, b string) int {

	// dp of i depends oppon i + 1 & i
	// so current row depends opon the next row
	nn := len(b)

	curr := make([]int, nn+1)
	next := make([]int, nn+1)
	// dp[i+1] mean here use the next && dp[i] mean we talk to the current one.

	// base case no need cause dp array already init with 0.
	for i := len(a) - 1; i >= 0; i-- {
		for j := len(b) - 1; j >= 0; j-- {

			ans := 0
			if a[i] == b[j] { //match condition
				ans = 1 + next[j+1]
			} else {
				// here the if a[i] != b[j] so we have two options
				// we can increment the i pointer and compare with whole j or
				// we can increment the j pointer and compare
				 with whole i. and store the max of both calls.
				ans = int(math.Max(float64(next[j]), float64(curr[j+1])))
			}
			curr[j] = ans

		}
		next = curr
	}
	fmt.Println("bottom up", next, curr)
	return next[0]
}

func LongestCommonSubsequence(text1, text2 string) int {

	return spaceOptimization(text1, text2)

	// can we do the space optimization. how to check?
	// it only happens when the current result depends on the previously 2 or 3 results.
}
