package array

import "log"

func Reverse(arr []int) {
	// TODO: Reverse the array.
	start := 0
	end := len(arr) - 1
	for start < end {
		log.Println("len", start, end)

		arr[start], arr[end] = arr[end], arr[start]
		start++
		end--
	}
}
