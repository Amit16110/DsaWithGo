package striing

import (
	"fmt"
	"strconv"
	"strings"
)

func CountAndSay(str string) {

	ans := CountAndSayHelper(str)
	fmt.Println("ans=>", ans)

}
func CountAndSayHelper(str string) string {
	if str == "" {
		return "1"
	}

	var result strings.Builder
	count := 1

	for i := 1; i < len(str); i++ {
		if str[i] == str[i-1] {
			count++
		} else {
			result.WriteString(strconv.Itoa(count))
			result.WriteByte(str[i-1])
			count = 1
		}

	}
	result.WriteString(strconv.Itoa(count))

	result.WriteByte(str[len(str)-1])
	return result.String()
}
