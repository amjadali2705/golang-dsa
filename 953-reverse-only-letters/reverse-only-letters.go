func reverseOnlyLetters(s string) string {
    runes := []rune(s)
	left, right := 0, len(runes)-1

	for left < right {
		for left < right && !unicode.IsLetter(runes[left]) {
			left++
		}

		for left < right && !unicode.IsLetter(runes[right]) {
			right--
		}

		if left < right {
			runes[left], runes[right] = runes[right], runes[left]
			left++
			right--
		}
	}

	return string(runes)
}