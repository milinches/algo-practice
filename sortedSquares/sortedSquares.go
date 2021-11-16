package main

import "sort"

func main() {
	d := []int{-7, -3, 2, 3, 11}
	sortedSquares(d)
}

func sortedSquares(nums []int) []int {
	squares := []int{}
	for _, value := range nums {
		values := value * *&value
		squares = append(squares, values)
	}
	sort.Ints(squares)
	return squares
}