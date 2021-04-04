/*
Given an input string "addc" string the first not repeating char
*/
package main

import (
	"fmt"
)

func main() {

	str := "aaaaabcdef"
	index := -1

	for i := 1; i < len(str); i++ {
		if str[i] != str[i-1] {
			fmt.Println("first non-repeating char is: ", string(str[i]))
			fmt.Println("first non-repeating char index is: ", i)
			index = i
			break
		}
	}

	fmt.Println("first non-repeating char index is: ", index)
}
