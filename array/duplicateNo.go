package array

import (
	"fmt"
)

func DuplicateNo(arr []int) int {
	//Require. n
	slow, fast := arr[0], arr[arr[0]]
	fmt.Println("slow, fast", slow, fast)
	for slow != fast {
		slow = arr[slow]
		fast = arr[arr[fast]]
		fmt.Println("slow, fast", slow, fast)
	}
	fast = 0

	for slow != fast {
		slow = arr[slow]
		fast = arr[fast]
		fmt.Println("slow, inside", slow, fast)
	}
	return slow
}

// find first smallest no. missing.
func MissingNumber(nums []int) int {
	//In range of n.
	//In this question already given the array must be in range of 0 -> n
	n := len(nums)
	var current, total int

	// for _, val := range nums {
	// 	if val <= n {
	// 		current += val
	// 	}
	// }
	// for i := 0; i < n+1; i++ {
	// 	total += i
	// 	fmt.Println("count", i, total)
	// }

	//UPATED OF OPTIMISE SOLUTION.

	for i := 0; i < n+1; i++ {
		total += i
		fmt.Println("show this", i, total)
		if i < n {
			current += nums[i]
			fmt.Println(current)
		}
	}
	return total - current
}

// This is some serious question but little bit a different varity.
