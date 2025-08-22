func addStrings(num1 string, num2 string) string {
    res := []byte{}
	i, j := len(num1)-1, len(num2)-1
	carry := 0

	for i >= 0 || j >= 0 || carry > 0 {
		sum := carry
		if i >= 0 {
			sum += int(num1[i] - '0')
			i--
		}
        
		if j >= 0 {
			sum += int(num2[j] - '0')
			j--
		}

		res = append(res, byte(sum%10)+'0')
		carry = sum / 10
	}

	for k, l := 0, len(res)-1; k < l; k, l = k+1, l-1 {
		res[k], res[l] = res[l], res[k]
	}

	return string(res)
}