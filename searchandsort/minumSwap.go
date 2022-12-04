package searchandsort

import (
	"fmt"
	"sort"
)

type Pair struct {
	val int
	ind int
}

func MiniSwap(arr []int) int {
	n := len(arr)
	// make a pair index array
	ans := make([]Pair, n)

	for i := 0; i < n; i++ {

		ans[i] = Pair{arr[i], i}

	}
	//Sort the array
	sort.Slice(ans, func(i, j int) bool {
		return ans[i].val < ans[j].val
	})
	count := 0
	fmt.Println("sort", ans)
	for i := 0; i < n; i++ {
		if ans[i].ind != i {
			count++
			ans[i], ans[ans[i].ind] = ans[ans[i].ind], ans[i]
			i--
		}
	}

	return count
}
