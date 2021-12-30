package main

import "fmt"

func TwoSum(nums []int, target int) []int {
	n := len(nums)
	for i := 0; i < n; i++ {
		for j := i; j < n; j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return []int{}
}

func main() {
	d := TwoSum([]int{1, 2, 3, 4, 5}, 3)
	fmt.Println(d)
}
