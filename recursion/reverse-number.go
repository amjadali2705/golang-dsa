package main

func ReverseDigits(n int) int {
	return reverseHelper(n, 0)
}

func reverseHelper(n, rev int) int {
	if n == 0 {
		return rev
	}

	rem := n % 10
	rev = rev * 10 + rem
	return reverseHelper(n / 10, rev)
}