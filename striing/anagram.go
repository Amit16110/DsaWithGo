package striing

func Anagram(s, t string) bool {
	// need to check t is anagram of s.
	if len(t) != len(s) {
		return false
	}
	hash := make(map[string]int)

	for _, v := range s {
		i := string(v)
		if hash[i] != 0 {

			hash[i] = hash[i] + 1
		} else {

			hash[i] = 1
		}
	}

	//check t string contain all the of hash if it not return false
	for _, v := range t {
		i := string(v)

		if hash[i] == 0 {
			return false
		} else {
			hash[i] = hash[i] - 1
		}
	}
	return true
}
