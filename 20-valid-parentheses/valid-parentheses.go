func isValid(s string) bool {
	bracketMap := map[rune]rune{
		')': '(',
		'}': '{',
		']': '[',
	}

	stack := []rune{}

	for _, char := range s {
		if openChar, isClosing := bracketMap[char]; isClosing {
			if len(stack) == 0 {
				return false
			}

			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1] 

			if top != openChar {
				return false
			}
		} else {
			stack = append(stack, char)
		}
	}

	return len(stack) == 0
}