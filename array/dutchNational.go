package array

//Sort an array of 0s, 1s and 2s

func DutchNational(arr []int) {
	low := 0
	mid := 0
	high := len(arr) - 1

	for mid <= high {

		if arr[mid] == 0 {
			// swap the ele with low.

			arr[low], arr[mid] = arr[mid], arr[low]
			low++
			mid++

		} else if arr[mid] == 1 {
			mid++
		} else {
			arr[mid], arr[high] = arr[high], arr[mid]
			high--
		}

	}

}
