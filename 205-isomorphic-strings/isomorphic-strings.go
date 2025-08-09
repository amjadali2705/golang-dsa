func isIsomorphic(s string, t string) bool {
    if len(s) != len(t) {
		return false
	}

	sToT := make(map[rune]rune)
	tToS := make(map[rune]rune)

	for i := 0; i < len(s); i++ {
		charS := rune(s[i])
		charT := rune(t[i])

		// Check mapping from s to t
		if val, ok := sToT[charS]; ok {
			if val != charT {
				return false
			}
		} else {
			sToT[charS] = charT
		}

		// Check mapping from t to s (to ensure one-to-one)
		if val, ok := tToS[charT]; ok {
			if val != charS {
				return false
			}
		} else {
			tToS[charT] = charS
		}
	}

	return true
}