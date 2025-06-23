func isPalindrome(s string) bool {
    nonAlphanumericRegex := regexp.MustCompile(`[^a-zA-Z0-9]+`)

	cleaned := nonAlphanumericRegex.ReplaceAllString(s, "") 
	cleaned = strings.ToLower(cleaned)                       

	left := 0
	right := len(cleaned) - 1

	for left < right {
		if cleaned[left] != cleaned[right] {
			return false
		}
		left++
		right--
	}
	return true
}