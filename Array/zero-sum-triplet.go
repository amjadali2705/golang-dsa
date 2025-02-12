package main

func ZeroSumTriplets(list []int) [][]int {
	res := make([][]int, 0)

	for i := 0; i < len(list); i++ {
		for j := i + 1; j < len(list); j++ {
			for k := j + 1; k < len(list); k++ {
				if list[i]+list[j]+list[k] == 0 {
					res = append(res, []int{list[i], list[j], list[k]})
				}
			}
		}
	}
	return res
}
