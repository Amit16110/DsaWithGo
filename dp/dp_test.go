package dp

import (
	"fmt"
	"testing"
)

func TestCoin(t *testing.T) {
	ans := coin([]int{1, 2, 5}, 7)

	fmt.Println("ans:", ans)
}
