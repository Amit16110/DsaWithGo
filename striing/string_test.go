package striing

import (
	"fmt"
	"testing"
)

func TestAnagram(t *testing.T) {
	arr := []string{"geeksquiz", "geeksforgeeks", "abcd", "forgeeksgeeks", "zuiqkeegs"}
	val := findAnagram(arr)
	fmt.Println("val", val)
}

func TestCountAndSay(t *testing.T) {
	CountAndSay("24")
}

func TestLS(t *testing.T) {
	no := longestSubstringWithoutRepeating("abcdacda")
	fmt.Println(" N", no)

}
func TestAtoi(t *testing.T) {
	no := myAtoi("abc")
	fmt.Println(no)

}
