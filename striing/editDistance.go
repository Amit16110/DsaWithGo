package striing

import "math"

func solveR(str1, str2 string, i, j int) int {
	// base case
	if i == len(str1) {
		return len(str2) - j
	}
	if j == len(str2) {
		return len(str1) - i
	}

	ans := 0

	if str1[i] == str2[j] {
		return solveR(str1, str2, i+1, j+1)
	} else {
		// insert
		in := 1 + solveR(str1, str2, i, j+1)
		// delete
		del := 1 + solveR(str1, str2, i+1, j)
		// replace
		rep := 1 + solveR(str1, str2, i+1, j+1)

		ans = int(math.Min(float64(in), math.Min(float64(del), float64(rep))))
	}
	return ans
}

func EditDistance(str1, str2 string) int {
	return solveR(str1, str2, 0, 0)
}
