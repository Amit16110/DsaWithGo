package tree

import "math"

func MaxheightOfTree(node *Node) int {
	if node == nil {
		return 0
	}

	left := MaxheightOfTree(node.left)
	right := MaxheightOfTree(node.right)

	ans := int(math.Max(float64(left), float64(right))) + 1
	return ans
}
