package striing

import "strings"

func GoalParser(s string) string {
	result := ""
	// for i := 0; i < len(s); i++ {

	// }
	result = strings.NewReplacer("()", "o", "(al)", "al").Replace(s)
	return result
}

func Interpreter(cmd string) string {
	// optimzed and simple solution.
	// ans := ""    always use tha var to declare the variable to speed up
	var ans string

	for i := 0; i < len(cmd); i++ {
		convStr := string(cmd[i])
		if convStr == "G" || convStr == "g" {
			ans += "G"
		}
		if convStr == "(" {
			i++
			if string(cmd[i]) == ")" {
				ans += "o"
			} else {
				ans += "al"
			}
		}
	}
	return ans
}
