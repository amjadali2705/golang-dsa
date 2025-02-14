package main

func InsertionSort(input []int) []int {
	for i := 1; i < len(input); i++ {
		j := i - 1
		key := input[i]

		for j >= 0 && input[j] > key {
			input[j+1] = input[j]
			j--
		}
		input[j+1] = key
	}
	return input
}
