func truncateSentence(s string, k int) string {
    words := strings.Split(s, " ")
	if k >= len(words) {
		return s 
	}
	truncatedWords := words[:k]
	return strings.Join(truncatedWords, " ")
}