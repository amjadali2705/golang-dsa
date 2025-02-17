package main

func LongestSubstringOfTwoChars(input string) string {
	if len(input) <= 1 {
		return ""
	}

	start, maxStart, maxLen := 0, 0, 0
	charCount := make(map[byte]int)

	for end := 0; end < len(input); end++ {
		charCount[input[end]]++

		for len(charCount) > 2 {
			charCount[input[start]]--
			if charCount[input[start]] == 0 {
				delete(charCount, input[start])
			}
			start++
		}

		if end-start+1 > maxLen {
			maxLen = end - start + 1
			maxStart = start
		}
	}
	return input[maxStart : maxStart+maxLen]
}
