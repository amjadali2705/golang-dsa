func repeatedSubstringPattern(s string) bool {

	if len(s) <= 1 {
		return false
	}
	
	temp := s + s
	
	modifiedTemp := temp[1 : len(temp)-1]
	
	return strings.Contains(modifiedTemp, s)
}