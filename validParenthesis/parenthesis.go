package main

import "fmt"

func main() {
	a := "()[]{}[]"
	fmt.Println(isValid(a))
}

func isValid(s string) bool {
	emptyValue := make([]rune, 0, len(s))
	mapping := map[rune]rune{')': '(', '}': '{', ']': '['}

	for _, value := range s {
		if value == '(' || value == '[' || value == '{' {
			emptyValue = append(emptyValue, value)
		} else {
			//if len(emptyValue) != 0 && emptyValue[len(emptyValue)-1] == mapping[value] {
			//	fmt.Println(string(emptyValue[len(emptyValue)-1]))
			//	emptyValue = emptyValue[:len(emptyValue)-1]
			//} else {
			//	return false
			//}
			if len(emptyValue) == 0 || mapping[value] != emptyValue[len(emptyValue)-1] {
				return false
			}
			emptyValue = emptyValue[:len(emptyValue)-1]
		}
	}
	return len(emptyValue) == 0
}
