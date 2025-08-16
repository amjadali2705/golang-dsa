func firstUniqChar(s string) int {
    charCount := make(map[rune]int)

	for _, char := range s {
		charCount[char]++
	}

	for i, char := range s {
		if charCount[char] == 1 {
			return i
		}
	}

	return -1
}