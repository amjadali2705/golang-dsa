package main

func BinarySearch(list []int, search int) int {
	start := 0
	end := len(list) - 1

	for start <= end {
		mid := start + (end - start) / 2

		if list[mid] == search {
			return mid
		} else if list[mid] < search {
			start = mid + 1
		} else {
			end = mid - 1
		}
	}
	return -1
}