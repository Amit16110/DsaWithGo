package striing

import "fmt"

// Brute force approch. n3
func LongestPalindromeSubstring(s string) int {
	n := len(s)
	// all substring of length 1 are palindrome,
	started := 0
	maxLength := 1

	for i := 0; i < n; i++ {
		for j := i; j < n; j++ {
			flag := 1

			// check palidrome
			for k := 0; k < (j-i+1)/2; k++ {
				if s[i+k] != s[j-k] {
					flag = 0
				}
			}
			if flag != 0 && (j-i+1) > maxLength {
				started = i
				maxLength = j - i + 1
			}
		}
	}
	fmt.Println(started)
	return maxLength
}

// Brute force but with little optimization without the DP.

func LongestPalindromeOptimze(s string) string {
	n := len(s)

	if n == 0 {
		return ""
	}

	var l, r, pl, pr int
	for r < n {
		// gobble up dup chars
		for r+1 < n && s[l] == s[r+1] {
			r++
		}
		// find size of this palindrome
		for l-1 >= 0 && r+1 < n && s[l-1] == s[r+1] {
			l--
			r++
		}
		if r-l > pr-pl {
			pl, pr = l, r
		}
		// reset to next mid point
		l = (l+r)/2 + 1
		r = l
	}
	return s[pl : pr+1]
}
