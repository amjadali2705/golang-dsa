func longestPalindrome(s string) int {
    counts := make(map[rune]int)
	for _, char := range s {
		counts[char]++
	}

	length := 0
	hasOdd := false

	for _, count := range counts {
		if count%2 == 0 {
			length += count
		} else {
			length += count - 1
			hasOdd = true
		}
	}

	if hasOdd {
		length++
	}

	return length
}