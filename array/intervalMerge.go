package array

func binarySearch(intervals [][]int, val int) int {
	left, right := 0, len(intervals)-1
	mid := right / 2

	for left <= right {

		if val < intervals[mid][0] {
			right = mid - 1
		} else if val > intervals[mid][1] {
			left = mid + 1
		} else {
			return mid
		}
		mid = (right-left)/2 + left
	}

	return left
}

func includeIntervalValue(interval []int, val int) bool {
	return val >= interval[0] && val <= interval[1]
}

func minInterval(startInterval, newInterval []int) []int {
	if startInterval[0] < newInterval[0] {
		newInterval[0] = startInterval[0]
	}

	return newInterval
}

func maxInterval(endInterval, newInterval []int) []int {
	if endInterval[1] > newInterval[1] {
		newInterval[1] = endInterval[1]
	}

	return newInterval
}

// best solution for the intervals, learn this algo at evening.

func insertInterval(intervals [][]int, newInterval []int) [][]int {
	var (
		startIndex = binarySearch(intervals, newInterval[0])
		endIndex   = binarySearch(intervals, newInterval[1])
		result     = make([][]int, 0, startIndex+len(intervals)-endIndex+1)
	)
	result = append(result, intervals[:startIndex]...)

	if startIndex < len(intervals) && includeIntervalValue(intervals[startIndex], newInterval[0]) {
		newInterval = minInterval(intervals[startIndex], newInterval)
	}
	if endIndex < len(intervals) && includeIntervalValue(intervals[endIndex], newInterval[1]) {
		newInterval = maxInterval(intervals[endIndex], newInterval)
	}

	result = append(result, newInterval)
	if endIndex < len(intervals) && includeIntervalValue(intervals[endIndex], newInterval[1]) {
		endIndex++
	}

	return append(result, intervals[endIndex:]...)
}
