package main

import (
	"strconv"
)

func main() {
	d := []int{555, 901, 482, 1771}
	println(findNumbers(d))
}

func findNumbers(nums []int) int {
	output := 0
	for _, value := range nums {
		if len(strconv.Itoa(value))%2 == 0 {
			output += 1
		}
	}
	return output
}
