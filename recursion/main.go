package main

import "fmt"

func main() {
	//reverse digits
	n := 12340
	fmt.Println(ReverseDigits(n))

	//check palindrome
	s := "aba"
	fmt.Println(IsPalindrome(s))

	//exponentiation
	fmt.Println(PowerOf(5, 4))

	//multiplication
	fmt.Println(Multiply(-3, -3))

	//expression operators
	fmt.Println(ExpressionOperators([]int{1, 2, 3, 4, 5, 6}, 5))
}
