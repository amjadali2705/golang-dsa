package main

func RotateKSteps(list []int, k int) []int {
	ReverseInPlace(list, 0, len(list)-k-1)
	ReverseInPlace(list, 0, len(list)-1)
	ReverseInPlace(list, 0, k-1)

	return list
}
