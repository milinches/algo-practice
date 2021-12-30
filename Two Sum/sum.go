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

func TwoSums(nums []int, target int) []int {
	makeMap := make(map[int]int)
	for index, value := range nums {
		if mapValue, okay := makeMap[target-value]; okay {
			return []int{mapValue, index}
		}
		makeMap[value] = index
	}
	return []int{}
}

func main() {
	d := TwoSum([]int{3, 2, 4}, 6)
	fmt.Println(d)
	s := TwoSums([]int{3, 2, 4}, 6)
	fmt.Println(s)
}
