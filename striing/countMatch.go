package striing

import "fmt"

// My solution take 60ms beat 50%
func CountMatch(items [][]string, ruleKey string, ruleValue string) int {
	data := []string{"type", "color", "name"} // get the index
	var indies int
	var count int
	for i := 0; i < len(data); i++ {
		if data[i] == ruleKey {
			indies = i
			break
		}
	}

	for i := 0; i < len(items); i++ {
		fmt.Println("items[i][indies]", items[i][indies])
		if items[i][indies] == ruleValue {

			count++
		}
	}
	return count
}

// optimse
func CountMatchOptimse(items [][]string, ruleKey string, ruleValue string) int {
	var ans int
	var index int

	switch ruleKey {
	case "type":
		index = 0
	case "color":
		index = 1
	case "name":
		index = 2
	}

	for _, v := range items {
		if v[index] == ruleValue {
			ans++
		}
	}
	return ans
}
