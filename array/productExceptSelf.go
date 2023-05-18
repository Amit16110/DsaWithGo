package array

func ProductExceptSelf(nums []int) []int {
	// base caeses
	n := len(nums)
	if n <= 1 {
		return nums
	}

	prodLeft := make([]int, n)
	// push 1 at the index 0
	prodLeft[0] = 1
	//  loop. Loop started at the 1
	for i := 1; i < n; i++ {
		prodLeft[i] = nums[i-1] * prodLeft[i-1]
	}

	prodRight := nums[n-1]

	for i := n - 2; i >= 0; i-- {
		prodLeft[i] *= prodRight
		prodRight *= nums[i]
	}

	return prodLeft
}
