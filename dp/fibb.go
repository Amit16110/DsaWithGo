package dp

import (
	"fmt"
	"log"
)

// fibbonacci => n is addition of previous two no's.

// Recursion without memorisation
// func FibbWithRec(n int) int {
// 	if n <= 1 {
// 		return n
// 	}
// 	return FibbWithRec(n-1) + FibbWithRec(n-2)
// }

func FibbWithDP(n int, dp []int) int {
	// fmt.Println("Dp inside==>", dp)
	if n <= 1 {
		return n
	}
	// Step 3. => check in value already stored or not.
	if dp[n] != -1 {
		return dp[n]
	}

	//Step 2 => store value of func in dp
	dp[n] = FibbWithDP(n-1, dp) + FibbWithDP(n-2, dp)
	return dp[n]

}
func FibbMain(n int) {
	log.Println("n=>", n)
	// Step 1 => create Dp array.
	dp := []int{}
	for i := 0; i <= n; i++ {
		dp = append(dp, -1)
	}

	log.Println("ans =>", FibbWithDP(n, dp))
}

func FibbTab(n int) {
	// STEP  0. Create Dp array.

	dp := make([]int, n+1)
	// STEP 1. Check base case.

	// And the base case is n == 0 => 0 and n == 1 => 1
	// So make them dp[0] = 0 & dp[1] = 1

	dp[0] = 0
	dp[1] = 1

	// STEP 3. convert the recursion in dp.
	for i := 2; i <= n; i++ {

		dp[i] = dp[i-1] + dp[i-2]
	}
	fmt.Println("Output", dp[n])
}

func FibbSO(n int) {
	// there is no need to store all the previous values.
	first := 0
	second := 1

	var result int
	for i := 2; i <= n; i++ {
		result = first + second
		first = second
		second = result
	}
	fmt.Println("Result", result)
}
