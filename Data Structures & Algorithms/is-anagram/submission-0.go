func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	hashMap := make(map[rune]int)

	for _, char := range s {
		hashMap[char] += 1
	}

	for _, char2 := range t {
		hashMap[char2] -= 1
	}

	for _, val := range hashMap {
		if val > 0 {
			return false
		}
	}

	return true
}
