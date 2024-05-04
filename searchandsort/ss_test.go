package searchandsort

import (
	"fmt"
	"testing"
)

func TestQuickSort(t *testing.T) {
	val := QuickSortStart([]int{1, 5, 10, 3, 1})
	fmt.Println(val)
}
