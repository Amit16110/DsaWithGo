package array

import (
	"fmt"
	"testing"
)

func TestMaxSumArray(t *testing.T) {
	// Kadane's algorithm
	array := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	maxSum := maxSubArray(array)
	fmt.Printf("max sum: %v", maxSum)

}

func TestExample(t *testing.T) {
	callingMyFunc()
}

func TestLongestPrefix(t *testing.T) {
	strs := []string{"aaa", "aaaabcadfdsafsafdsjk", "aaaaaaaaaabbacasdfsadfdaaaaaa", "throught", "theeeeeeeeee"}
	ans := longestCommonPrefix(strs)
	fmt.Println("ans", ans)

}

func TestCountAnagram(t *testing.T) {
	str := "forxxorfxdofr"
	pat := "fxd"

	val := findAnagrams(str, pat)
	fmt.Println("val", val)
}
func TestIntervalMerge2(t *testing.T) {
	intervals := [][]int{{1, 2}, {3, 5}, {6, 7}, {8, 10}, {12, 16}}
	newInterval := []int{4, 8}

	val := insertInterval(intervals, newInterval)
	fmt.Println("val", val)
}

func TestTwoSum(t *testing.T) {
	arr := []int{2, 6, 5, 8, 11}
	target := 50

	result := twoSum(arr, target)

	fmt.Println("result", result)
}
