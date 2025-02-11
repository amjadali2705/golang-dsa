package main

func ReverseInPlace(list []int, start, end int) {

	for start < end {
		list[start], list[end] = list[end], list[start]
		start++
		end--
	}

}
