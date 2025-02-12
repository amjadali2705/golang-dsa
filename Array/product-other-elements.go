package main

func ProductOfAllOtherElements(list []int) []int {
	res := make([]int, 0)

	for i := 0; i < len(list); i++ {
		prod := 1
		for j := 0; j < len(list); j++ {
			if i != j {
				prod *= list[j]
			}
		}
		res = append(res, prod)
	}
	return res
}