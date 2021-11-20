package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(findMedianSortedArrays([]int{1, 2}, []int{3, 4}))
}

func findMedianSortedArrays(nums1 []int, num2 []int) float64 {
	keys := make(map[int]bool)
	list := []int{}
	num := append(nums1, num2...)
	sort.Ints(num)
	addNumbers := 0

	for _, entry := range num {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}

	for _, number := range list {
		addNumbers = addNumbers + number
		
	}
	
	findMedian := float64(addNumbers) / float64(len(list))

	return findMedian	
}