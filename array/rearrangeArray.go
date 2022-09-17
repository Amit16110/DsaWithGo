package array

import (
	"fmt"
)

// Rearrange the array in O(1) space.

func RearrangeArray(arr []int) {
	// Idea. take every negative ele to the right side of array and after place them alternatively

	n := len(arr)
	i, j := 0, n-1
	for i < j {
		// First loop for detect the first negative
		for i <= n-1 && arr[i] > 0 {
			i += 1
		}
		//Second loop for detect the first positive
		for j >= 0 && arr[j] < 0 {
			j -= 1
		}

		if i < j {
			swap(arr, i, j)
		}
	}
	if i == 0 && i == n {
		return
	}

	k := 0
	//swap the right side negitive to left side positive.
	for k < n && i < n {
		swap(arr, k, i)
		i = i + 1
		k = k + 2
	}
	fmt.Println("array==>", arr)
}

func swap(arr []int, i, j int) {

	arr[i], arr[j] = arr[j], arr[i]
}
