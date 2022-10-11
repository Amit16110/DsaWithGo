package dp

import (
	"fmt"
	"math"
)

func CutRoadInSegment(n, x, y, z int) int {
	ans := recCutRoad(n, x, y, z)
	fmt.Println("ans", ans)
	if ans < 0 {
		return 0
	} else {
		return ans
	}
}

func recCutRoad(n, x, y, z int) int {
	if n == 0 {
		return 0
	}
	if n < 0 {
		return math.MinInt
	}

	a := recCutRoad(n-x, x, y, z) + 1
	b := recCutRoad(n-y, x, y, z) + 1
	c := recCutRoad(n-z, x, y, z) + 1

	ans := int(math.Max(float64(a), math.Max(float64(b), float64(c))))
	return ans
}
