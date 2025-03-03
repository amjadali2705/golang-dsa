package main

func PowerOf(x, n int) int {
	res := 1

	if n != 0 {
		res = x * PowerOf(x, n-1)
	}
	
	return res
}