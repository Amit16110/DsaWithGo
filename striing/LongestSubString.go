package striing

import "math"

// longest substring without repeating the character.

func longestSubstringWithoutRepeating(str string) int {
	var start, result int

	charIndexMap := make(map[byte]int)

	for end := 0; end < len(str); end++ {
		if duplicateInd, ok := charIndexMap[str[end]]; ok {
			result = int(math.Max(float64(result), float64(end-start)))

			for i := start; i <= duplicateInd; i++ {
				delete(charIndexMap, str[i])
			}

			// Slide the window since we have found a duplicate character
			start = duplicateInd + 1
		}
		// Add the current character to hashmap
		charIndexMap[str[end]] = end

	}
	result = int(math.Max(float64(result), float64(len(str)-start)))
	return result
}
