package array

// ! NOT SOLVED YET BUT VERY IMPORTANT CONCEPT.
func sameprefixWithLongest(arr []string) {
	// j := 1
	for i := 0; i < len(arr); i++ {
		var prefix byte
		for j := 0; j < len(arr[i]); j++ {
			for k := 1; k < len(arr); k++ {
				if arr[i][j] == arr[k][j] {
					prefix = prefix + arr[i][j]
				}
			}

			//

		}

	}
}

// func longestCommonPrefix(strs []string) string {
// 	if len(strs) == 0 {
// 		return ""
// 	}
// 	prefix := strs[0]
// 	for i := 1; i < len(strs); i++ {
// 		fmt.Println("str", strings.HasPrefix(strs[i], prefix))
// 		for !strings.HasPrefix(strs[i], prefix) {
// 			prefix = prefix[:len(prefix)-1]
// 			if len(prefix) == 0 {
// 				return ""
// 			}
// 		}
// 	}
// 	return prefix
// }

type TrieNode struct {
	children map[rune]*TrieNode
	word     string
}

func insert(node *TrieNode, word string) {
	for _, r := range word {
		child, exists := node.children[r]
		if !exists {
			child = &TrieNode{children: make(map[rune]*TrieNode)}
			node.children[r] = child
		}
		node = child
	}
	node.word = word
}

func findLongestCommonPrefix(node *TrieNode) string {
	var longestPrefix string
	for _, child := range node.children {
		prefix := findLongestCommonPrefix(child)
		if len(prefix) > len(longestPrefix) {
			longestPrefix = prefix
		}
	}
	if node.word != "" && len(node.word) > len(longestPrefix) {
		longestPrefix = node.word
	}
	return longestPrefix
}

func longestCommonPrefix(strings []string) []string {
	root := &TrieNode{children: make(map[rune]*TrieNode)}
	for _, s := range strings {
		insert(root, s)
	}
	longestPrefix := findLongestCommonPrefix(root)
	result := []string{}
	for _, s := range strings {
		if len(s) >= len(longestPrefix) && s[:len(longestPrefix)] == longestPrefix {
			result = append(result, s)
		}
	}
	return result
}
