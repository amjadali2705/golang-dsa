package main

func MergeSort(list []int) []int {
	if len(list) <= 1 {
		return list
	}

	left, right := divide(list)
	return merge(left, right)
}

func divide(list []int) ([]int, []int) {
	mid := len(list) / 2
	left := MergeSort(list[:mid])
	right := MergeSort(list[mid:])
	return left, right
}

func merge(left, right []int) []int {
	res := []int{}
	i, j := 0, 0

	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			res = append(res, left[i])
			i++
		} else {
			res = append(res, right[j])
			j++
		}
	}

	if i >= len(left) {
		res = append(res, right[j:]...)
	} else {
		res = append(res, left[i:]...)
	}
	return res
}