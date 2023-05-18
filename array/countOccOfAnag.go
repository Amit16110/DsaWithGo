package array

import "fmt"

// This method is half required.
func CountOfAnagram(s, p string) {
	m := make(map[byte]int)

	// ans := 0
	count := 0

	left, right := 0, 0

	for i := 0; i < len(p); i++ {
		m[p[i]]++
		count = len(m)
	}

	for right < len(s) {

		if right-left+1 < len(p) {
			if val, ok := m[s[right]]; ok {
				m[s[right]]--

				if val == 0 {
					count--
				}
			}
			right++
		}

	}
}

func findAnagrams(s string, p string) []int {
	pFreq := [26]int{}
	for _, c := range p {
		pFreq[c-'a']++
	}
	fmt.Println("pfreq", pFreq)
	var indices []int
	freq := [26]int{}
	for i := 0; i < len(s); i++ {
		if i >= len(p) {
			freq[s[i-len(p)]-'a']--
		}
		freq[s[i]-'a']++
		fmt.Println("freq", freq)

		if freq == pFreq {
			indices = append(indices, i-(len(p)-1))
		}
	}
	return indices
}
