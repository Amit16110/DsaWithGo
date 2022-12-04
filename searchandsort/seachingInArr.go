package searchandsort

import (
	"fmt"
	"math"
)

// searching in an array where adjacent diff by most k.

func SearchInArr(arr []int, x, k int) int {
	n := len(arr)

	// Traverse the given array starting from the leftmost ele.
	start := 0

	for start < n {
		// if x is found at index i
		if x == arr[start] {
			return start
		}
		// Main logic	:=> jump the difference between	current
		// - arr ele and x divided by k.
		// we use max here to make sure that i moves at - least one step ahead;
		fmt.Println(">>>>>>", float64(arr[start]-x)/float64(k), arr[start])
		start = start + int(math.Max(1, math.Abs(float64(arr[start]-x)/float64(k))))
		fmt.Println("start", start)
	}

	return -1
}
