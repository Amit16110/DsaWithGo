package searchandsort

func SearchInSortAndRotated(nums []int, target int) int {
	// TODO: working on this
	left, right := 0, len(nums)-1
	for left <= right {
		mid := (left + right) >> 1

		if nums[mid] == target {
			return mid
		}

		if nums[left] <= nums[right] {
			if target > nums[mid] {
				left = mid + 1
			} else {
				right = mid - 1
			}
		} else {
			if target > nums[mid] {
				if nums[mid] >= nums[left] {
					left = mid + 1
				} else {
					if target <= nums[right] {
						left = mid + 1
					} else {
						right = mid - 1
					}
				}
			} else {
				if nums[mid] >= nums[left] {
					if target >= nums[left] {
						right = mid - 1
					} else {
						left = mid + 1
					}
				} else {
					right = mid - 1
				}
			}
		}
	}
	return -1
}
