func reverseVowels(s string) string {
    vowels := "aeiouAEIOU"
	sList := []rune(s)
	left, right := 0, len(sList)-1

	for left < right {
		for left < right && !strings.ContainsRune(vowels, sList[left]) {
			left++
		}

		for left < right && !strings.ContainsRune(vowels, sList[right]) {
			right--
		}

		if left < right {
			sList[left], sList[right] = sList[right], sList[left]
			left++
			right--
		}
	}

	return string(sList)
}