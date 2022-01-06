package main

import (
	"fmt"
	"testing"
)

// All available "test cases," you can add more if you want to.
var testCases = []string{
	"{}[[]]", "{[](()[])}", "(]", "([))]", "{()}",
}

// Function TestName is where the main logic lies.
//
// It loops through the testCases, and then checks if the parenthesis is Valid or not.
func TestName(t *testing.T) {
	for _, tests := range testCases {
		testInput := isValid(tests)

		// if parenthesis is not valid, throw new Error("Failed test ❌")
		if testInput != true {
			t.Error("Failed test ❌")
		} else {
			fmt.Println("Test Passed ✔") // Test passed
		}
	}
}
