package main

import "fmt"

func main() {

	list := []int{1, 2, 3, 4, 5, 6}
	start := 0
	end := 4

	ReverseInPlace(list, start, end)
	fmt.Println(list)

	//Sum of numbers in two slices
	num1 := []int{2, 9}
	num2 := []int{9, 9, 9}
	fmt.Println(AddTwoNumbers(num1, num2))
}
