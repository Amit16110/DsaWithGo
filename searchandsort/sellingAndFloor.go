package searchandsort

import "fmt"

func CeilingOfNo(arr []int, target int) int {
	// return a number that are greater than or equal to target element(>=target)

	start, end := 0, len(arr)-1
	fmt.Println("start", start, end)
	for start <= end {
		mid := start + (end-start)/2
		fmt.Println("mid", arr[mid],mid)
		if arr[mid] == target {
			return mid
		} else if target > arr[mid] {
			start = mid + 1
		} else {
			end = mid - 1
		}

	}

	return start
}
