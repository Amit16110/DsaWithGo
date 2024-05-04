package striing

import (
	"math"
	"strings"
)

func myAtoi(str string) int {
	// trim the string
	s := strings.Trim(str, " ")

	n := len(str)

	if n == 0 {
		return 0
	}

	var start int

	// handling numbers with +/- sign
	signMultiplex := 1

	if s[0] == '-' {
		// handling for the negative numbers
		signMultiplex = -1
		start = 1
	} else if s[0] == '+' {
		start = 1
	}

	// ASCII of '0' = 48
	// s[i] - '0' gives the actual number as an integer

	var res int
	for i := start; i < len(s); i++ {
		// Handling for non numeric values

		if !(s[i] >= '0' && s[i] <= '9') {
			return res * signMultiplex
		}
		// cal the result for current iteration
		res = res*10 + int(s[i]-'0')

		// Check if result exceeds MinInt32 or MaxInt32

		if res*signMultiplex <= math.MinInt32 {
			return math.MinInt32
		} else if res*signMultiplex >= math.MaxInt32 {
			return math.MaxInt32
		}
	}
	return res * signMultiplex
}
