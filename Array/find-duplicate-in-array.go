package main 

func FindDuplicate(list []int) int {
	visited := make(map[int]bool)

	for i :=  0; i < len(list); i++ {
		if visited[list[i]] {
			return list[i]
		} else {
			visited[list[i]] = true
		}
	}
	return -1
}