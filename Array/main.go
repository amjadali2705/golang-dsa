package main

import "fmt"

func main() {
	
	list := []int{1, 2, 3, 4, 5, 6}
	start := 0
	end := 4

	ReverseInPlace(list, start, end)
	fmt.Println(list)
}
