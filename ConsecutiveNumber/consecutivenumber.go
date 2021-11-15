package main

// Given a binary array nums, return the maximum number of consecutive 1's in the array.

func main() {
	d := []int{1, 1, 0, 1, 1, 1}
	println(findMaxConsecutiveNumber(d))
}

func findMaxConsecutiveNumber(nums []int) int {
	max := 0
	count := 0
	for _, value := range nums {
		if value == 1 {
			count += 1
			if max < count {
				max = count
			}
		} else {
			count = 0
		}
	}
	return max
}
