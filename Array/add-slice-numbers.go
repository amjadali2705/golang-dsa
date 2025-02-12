package main

func AddTwoNumbers(num1, num2 []int) []int {
	num1, num2 = convertEqLen(num1, num2)
	carry := false

	for i := len(num1) - 1; i >= 0; i-- {
		num1[i] += num2[i]
		if carry {
			num1[i]++
			carry = false
		}

		if num1[i] >= 10 {
			num1[i] -= 10
			carry = true
		}
	}

	if carry {
		num1 = append([]int{1}, num1...)
	}
	return num1
}

func convertEqLen(num1, num2 []int) ([]int, []int) {
	diff := absDiff(len(num1), len(num2))
	filledWithZeroes := make([]int, diff)

	if len(num1) > len(num2) {
		num2 = append(filledWithZeroes, num2...)
	} else if len(num2) > len(num1) {
		num1 = append(filledWithZeroes, num1...)
	}

	return num1, num2
}

func absDiff(a, b int) int {
	if a > b {
		return a - b
	} else {
		return b - a
	}
}
