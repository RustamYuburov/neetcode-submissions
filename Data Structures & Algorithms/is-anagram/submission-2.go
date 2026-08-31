func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	hash := make(map[rune]int)

	for _, char := range s {
		hash[char] += 1
	}
	fmt.Println(hash)
	for _, char2 := range t {
		hash[char2] -= 1
	}
	fmt.Println(hash)
	for _, val := range hash {
		if val != 0 {
			return false
		}
	}

	return true
}
