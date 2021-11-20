package main

func main() {
	d := []int{1,0,2,3,0,4,5,0}
	duplicateZeros(d)
}

func duplicateZeros(nums []int) {
	newArr := []int{}
	for _, value := range nums {
		if nums[value] == 0 {
			newArr = append(newArr, value)
		}
	}
	println(*(&newArr))
}