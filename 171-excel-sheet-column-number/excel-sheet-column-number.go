func titleToNumber(columnTitle string) int {
    result := 0
	for i :=0; i < len(columnTitle); i++ {
		charValue := int(columnTitle[i] - 'A' + 1)
		result = result*26 + charValue
	}
	return result
}