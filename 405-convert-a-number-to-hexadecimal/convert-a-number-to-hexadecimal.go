func toHex(num int) string {
    if num == 0 {
		return "0"
	}

	hexChars := "0123456789abcdef"
	result := ""

	unsignedNum := uint32(num)

	for unsignedNum > 0 {
		digitValue := unsignedNum % 16
		result = string(hexChars[digitValue]) + result
		unsignedNum /= 16
	}

	return result
}