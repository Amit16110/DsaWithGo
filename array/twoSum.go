package array

func twoSum(arr []int, target int) []int {
	m := make(map[int]int)
	ans := make([]int, 2)

	for i, v := range arr {

		needed := target - v

		if m[needed] != 0 {
			ans[0] = m[needed]
			ans[1] = i
			return ans
		}
		m[v] = i
	}
	return ans
}
