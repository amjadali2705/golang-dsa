package main

import "fmt"

func main() {
	//binary search
	fmt.Println(BinarySearch([]int{1, 2, 3, 4, 5, 9}, 4))

	//square root
	fmt.Println(SquareRoot(5, 3))

	//merge sort
	fmt.Println(MergeSort([]int{-1, 3, 2, 0, 4}))

	//quick sort
	fmt.Println(QuickSort([]int{-1, 3, 2, 0, 4}))
}