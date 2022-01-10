package main

import (
	"fmt"
	"math"
	"strings"
)

func main() {
	var inputs = []string{
		"have a nice day",
		"stay cavy",
		"feed the dog",
		"chill out",
	}
	for _, outputs := range inputs {
		fmt.Println(encryption(outputs))
	}
}

func encryption(s string) string {
	// Remove all whitespaces from the string
	trimString := strings.ReplaceAll(s, " ", "")
	strLenght := len(trimString)
	findSqrt := math.Sqrt(float64(strLenght))
	rows := math.Floor(findSqrt)
	columns := math.Ceil(findSqrt)

	if int(rows)*int(columns) < strLenght {
		rows++
	}

	var output string = ""

	for i := 0; i < int(columns); i++ {
		for j := 0; j < int(rows); j++ {
			if j*int(columns)+i < strLenght {
				output += string(trimString[j*int(columns)+i])
			}
		}
		output += " "
	}
	return output
}
