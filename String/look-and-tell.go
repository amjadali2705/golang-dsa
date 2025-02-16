package main

import "strconv"

func LookAndTell(depth int) []string {
	if depth < 0 {
		return []string{"-1"}
	}

	result := []string{"1"}
	for i := 1; i < depth; i++ {
		result = append(result, generateNext(result[i-1]))
	}
	return result
}

func generateNext(prev string) string {
	count := 1
	next := ""

	for i := 1; i < len(prev); i++ {
		if prev[i] == prev[i-1] {
			count++
		} else {
			next += strconv.Itoa(count) + string(prev[i-1])
			count = 1
		}
	}
	next += strconv.Itoa(count) + string(prev[len(prev)-1])
	return next
}
