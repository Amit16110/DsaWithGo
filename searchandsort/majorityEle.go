package searchandsort

import "fmt"

func MajorityFirstType(arr []int) int {
	// Boyer-moore voting algo.

	var count, majorityEle int

	for i := 0; i < len(arr); i++ {
		if count == 0 {
			majorityEle = arr[i]
		}
		if majorityEle == arr[i] {
			count += 1
		} else {
			count -= 1
		}

		fmt.Println("count", count, arr[i])
		if count > len(arr)/2 {
			return majorityEle
		}
	}
	return -1
}
func MajorityElementTwo(nums []int) []int {
	var count, majorityElements int

	n := len(nums)

	for _, v := range nums {
		if count == 0 {
			majorityElements = v
		}
		if majorityElements == v {
			count += 1
		} else {
			count -= 1
		}
		fmt.Println("====>", count, n)
		if count > n/3 {
			return []int{majorityElements}
		}
	}
	return nums
}
