func licenseKeyFormatting(s string, k int) string {
	cleanedStr := strings.ToUpper(strings.ReplaceAll(s, "-", ""))
	n := len(cleanedStr)

	if n == 0 {
		return ""
	}

	var result strings.Builder
	firstGroupLen := n % k
	if firstGroupLen == 0 {
		firstGroupLen = k
	}

	if firstGroupLen > 0 && n > 0 {
		result.WriteString(cleanedStr[:firstGroupLen])
	}

	for i := firstGroupLen; i < n; i += k {
		result.WriteString("-")
		result.WriteString(cleanedStr[i : i+k])
	}

	return result.String()
}