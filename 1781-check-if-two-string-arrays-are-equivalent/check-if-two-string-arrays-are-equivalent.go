func arrayStringsAreEqual(word1 []string, word2 []string) bool {
    s1 := strings.Join(word1, "")

	s2 := strings.Join(word2, "")

	return s1 == s2
}