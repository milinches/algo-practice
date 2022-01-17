package main

import (
	"fmt"
	"math"
	"sort"
)

// We are going to search for a particular name in an Array of names using binary search.

func main() {
	words := []string{"Hello", "World", "Where", "Is", "The", "Way", "To", "The", "Available", "Bathroom", "I", "need",}
	fmt.Println(binarySearch(words, "Is"))
}

func binarySearch(searchArray []string, searchWord string) bool {
	// To acheive a binary sort, the string must be sorted.
	sort.Strings(searchArray)
	min := 0
	max := len(searchArray) - 1

	for {
		splitting := math.Floor(float64(min+max) / 2)
		convSplitting := int(splitting)

		if max < min {
			return false
		}

		if searchArray[convSplitting] == searchWord {
			return true
		}

		if searchArray[convSplitting] < searchWord {
			min = convSplitting + 1
		} else {
			max = convSplitting - 1
		}
	}
}