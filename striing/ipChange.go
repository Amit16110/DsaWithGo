package striing

import (
	"strings"
)

func IpChange(address string) string {
	n := len(address)
	if n < 1 {
		return " "
	}

	result := ""
	for i := 0; i < n; i++ {
		strChange := string(address[i])
		if strChange == "." {
			result = result + "[.]"
		} else {
			result = result + strChange
		}
	}
	return result
}

func IpChangeOptimse(address string) string {
	return strings.ReplaceAll(address, ".", "[.]")
}
