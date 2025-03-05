package main

func ExpressionOperators(list []int, target int) string {
	if len(list) == 0 {
		return ""
	}

	return findOperator(1, list[0], list, target, "")
}

func findOperator(index int, currentSum int, list []int, targetsum int, path string) string {
	if index == len(list) {
		if currentSum == targetsum {
			return path
		}
		return ""
	}

	plusResult := findOperator(index+1, currentSum+list[index], list, targetsum, path+"+")
	if plusResult != "" {
		return plusResult
	}

	minusResult := findOperator(index+1, currentSum-list[index], list, targetsum, path+"-")
	return minusResult
}