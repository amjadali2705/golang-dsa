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
}
