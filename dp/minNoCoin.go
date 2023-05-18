package dp

import (
	"math"
)

func coinRec(num []int, x int) int {
	if x == 0 {
		return 0
	}
	if x < 0 {
		return math.MaxInt
	}

	mini := math.MaxInt

	for i := 0; i < len(num); i++ {
		ans := coinRec(num, x-num[i])

		if ans != math.MaxInt {
			// 1 is used for the count of the coin
			mini = int(math.Min(float64(mini), float64(1+ans)))
		}
	}

	return mini
}

func coinMem(num, dp []int, x int) int {
	if x == 0 {
		return 0
	}
	if x < 0 {
		return math.MaxInt
	}
	// step 3: add a base case if call or index is found return the value.
	if dp[x] != 0 {
		return dp[x]
	}

	mini := math.MaxInt

	for i := 0; i < len(num); i++ {
		ans := coinMem(num, dp, x-num[i])

		if ans != math.MaxInt {
			// 1 is used for the count of the coin
			mini = int(math.Min(float64(mini), float64(1+ans)))
		}
	}
	// step 2: store the return value in the dp array.
	dp[x] = mini
	return mini
}

// Bottom up approch.
func coinTabulation(num []int, x int) int {
	// step 1: create a dp array
	dp := make([]int, x+1)

	dp[0] = 0

	for i := 1; i <= x; i++ {
		//  for any no.of index/x we require the all coins
		for j := 0; j <= len(num); j++ {
			dp[i] = int(math.Min(float64(dp[i]), float64(dp[i-num[j]])))
		}
	}
	return dp[x]
}

func coin(num []int, x int) int {
	// ans := coinRec(num, x)
	// if ans == math.MaxInt {
	// 	return -1
	// }
	// return ans

	// change recursive into memorisation
	// step 1
	dp := make([]int, x+1)
	ans := coinMem(num, dp, x)
	if ans == math.MaxInt {
		return -1
	}
	return ans
}
