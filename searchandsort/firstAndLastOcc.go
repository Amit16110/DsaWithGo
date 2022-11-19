package searchandsort

import "fmt"

func FirstAndLastOcc(arr []int, target int) []int {
	start := find(arr, target, true)
	fmt.Println("data show here")
	if start == -1 {
		return []int{-1, -1}
	}
	end := find(arr, target, false)

	return []int{start, end}

}

// write a Bs to search the target ele.

func find(arr []int, t int, start bool) int {
	n := len(arr)
	left, right := 0, n-1

	ans := -1

	for left <= right {
		mid := left + (right-left)/2

		if arr[mid] < t {
			left = mid + 1
		} else if arr[mid] > t {
			right = mid - 1
		} else {
			ans = mid

			if start {
				right = mid - 1
			} else {
				left = mid + 1
			}

		}
	}
	return ans
}
