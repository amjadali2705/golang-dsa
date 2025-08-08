func convertToTitle(columnNumber int) string {
    result := ""
	for columnNumber > 0 {
		columnNumber--
		rem := columnNumber % 26
		result = string('A'+rem) + result
		columnNumber /= 26
	}
	return result
}