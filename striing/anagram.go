package striing

import (
	"fmt"
	"sort"
	"strings"
)

func Anagram(s, t string) bool {
	// need to check t is anagram of s.
	if len(t) != len(s) {
		return false
	}
	hash := make(map[rune]int)

	for _, v := range s {

		hash[v]++
		// if hash[i] != 0 {

		// 	hash[i] = hash[i] + 1
		// } else {

		// 	hash[i] = 1
		// }
	}
	//check t string contain all the of hash if it not return false
	for _, v := range t {

		if hash[v] == 0 {
			return false
		} else {
			hash[v]--
			// hash[i] = hash[i] - 1
		}
	}
	return true
}

// tc => n^2
func anagramTest(str []string) {
	ans := []string{}
	// for i := 0; i < len(str); i++ {
	// 	for j := i + 1; j < len(str); j++ {
	// 		if i != j {
	// 			val := Anagram(str[i], str[j])
	// 			if val {
	// 				ans = append(ans, str[i], str[j])
	// 			}
	// 		}
	// 	}
	// }  update code
	for i, s1 := range str {
		for _, s2 := range str[i+1:] {
			if Anagram(s1, s2) {
				ans = append(ans, s1, s2)
			}
		}
	}
	fmt.Println("ans:", ans)
}

// Optimised version => n * nlogn and space complexity is n*k
func findAnagram(str []string) [][]string {
	groups := make(map[string][]string)

	for _, s := range str {
		sorted := sortStrings(s)
		groups[sorted] = append(groups[sorted], s)
	}
	result := make([][]string, 0, len(groups))
	for _, v := range groups {
		if len(v) > 1 {

			result = append(result, v)
		}
	}
	return result
}
func sortStrings(str string) string {
	s := strings.Split(str, "")
	sort.Strings(s)
	return strings.Join(s, "")
}
