package main

import "fmt"

func main() {
	d := []int{1, 2, 3, 4, 5, 6}
	fmt.Println(deleteElement(d))
}

func deleteElement(nums []int) []int {
	nums = nums[:len(nums) - 1]
	return nums
}