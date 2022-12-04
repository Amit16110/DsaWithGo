package searchandsort

import (
	"fmt"
	"math"
)

func isPossible(arr []int, n, m, curr_min int) bool {

	studentRequired := 1
	curr_sum := 0

	// iterate over all books
	for i := 0; i < n; i++ {
		curr_sum += arr[i]
		if curr_sum > curr_min {
			studentRequired++ //update student
			curr_sum++
		}
	}
	return studentRequired <= m
}
func FindPages(arr []int, m int) int {
	sum := 0
	n := len(arr)
	if n < m {
		return -1
	}

	// count total no. of pages.
	for _, v := range arr {
		sum += v
	}
	start, end := arr[n-1], sum
	result := math.MaxInt
	for start <= end {
		mid := start + (end-start)>>1

		fmt.Println("mid", mid)

		if isPossible(arr, n, m, mid) {
			//update result to current distribution
			// as it's the best we have found till now.
			result = mid
			end = mid - 1
		} else {
			// if not possible, means pages should be
			// increased ,so update start = mid + 1
			start = mid + 1
		}
	}
	return result
}
