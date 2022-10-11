package striing

// Check if the given strings are rotations of each other or not

func StringRotationCheck(str1, str2 string) bool {

	//base case
	if len(str1) != len(str2) {
		return false
	}

	// store the first occurences of the first character of str 1.
	indexes := []int{}
	size := len(str1)

	for i := 0; i < size; i++ {
		if str2[i] == str1[0] {
			indexes = append(indexes, i)
		}
	}
	isRotation := false

	for _, val := range indexes {

		isRotation = checkString(str1, str2, val, size)

		// if isRotation {
		// 	break
		// }
	}
	// if isRotation == true {
	// 	return true
	// }
	return isRotation
}
func checkString(str1, str2 string, indexFound, size int) bool {
	//check whether the character is equal or not
	for i := 0; i < size; i++ {
		if str1[i] != str2[(indexFound+i)%size] {
			return false
		}
	}
	return true
}
