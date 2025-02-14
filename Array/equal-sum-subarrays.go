package main

func EqualSubArrays(list []int) [][]int {
	res := make([][]int, 0)

	if len(list) < 2 {
		return res
	}

	splitPoint := findSplitPoint(list)
	if splitPoint == -1 || splitPoint == len(list) {
		return res
	}

	res = append(res, list[0:splitPoint])
	res = append(res, list[splitPoint:])

	return res
}

func findSplitPoint(list []int) int {
	leftSum := 0

	for i := 0; i < len(list); i++ {
		leftSum += list[i]

		rightSum := 0
		for j := i+1; j < len(list); j++ {
			rightSum += list[j]
		}

		if leftSum == rightSum {
			return i + 1
		}
	}
	return -1
}
