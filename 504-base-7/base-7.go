func convertToBase7(num int) string {
    if num == 0 {
		return "0"
	}

	isNegative := false
	if num < 0 {
		isNegative = true
		num = -num
	}

	result := ""
	for num > 0 {
		remainder := num % 7
		result = strconv.Itoa(remainder) + result
		num /= 7
	}

	if isNegative {
		return "-" + result
	}

	return result
}