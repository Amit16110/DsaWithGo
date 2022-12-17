package array

// Most Optimised solutions.
func Trapping(arr []int) int {
	n := len(arr)
	left, right := 0, n-1
	res := 0
	maxLeft, maxRight := 0, 0

	for left <= right {
		if arr[left] <= arr[right] {
			if arr[left] >= maxLeft {
				maxLeft = arr[left]
			} else {
				res += maxLeft - arr[left]
			}
			left++
		} else {
			if arr[right] >= maxRight {
				maxRight = arr[right]
			} else {
				res += maxRight - arr[right]
			}
			right--
		}
	}
	return res
}
