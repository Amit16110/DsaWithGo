package striing

import "fmt"

func Suffle(letter string, indexies []int) string {
	//1 map the data with string
	strIndex := make(map[int]string)

	for i := 0; i < len(letter); i++ {
		arrVal := indexies[i]
		strVal := string(letter[i])
		strIndex[arrVal] = strVal
	}
	result := ""
	fmt.Println("string value", strIndex)
	for i := 0; i < len(letter); i++ {
		result = result + strIndex[i]
	}
	return result
}

// optimeze approch.

func SuffleOptime(s string, indices []int) string {
	//1. string are made by the bytes so take the byte of array
	// is good for the string questions.
	bytes := make([]byte, len(indices))

	for i := 0; i < len(indices); i++ {
		bytes[indices[i]] = s[i]

	}

	str := fmt.Sprintf("%s", bytes)

	return str
}
