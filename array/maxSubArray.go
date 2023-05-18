package array

import (
	"math"
)

// kadane's algorithm

func maxSubArray(arr []int) int {
	sum := 0
	max := math.MinInt

	for _, v := range arr {
		sum += v
		max = int(math.Max(float64(sum), float64(max)))
		if sum < 0 {
			sum = 0
		}
	}
	// strconv.Atoi()
	return max
}
