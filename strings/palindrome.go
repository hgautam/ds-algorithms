/*
Checks if a string is a palindrome
input "abcdcba" will return in true
time complexity: o(n)
space complexity: o(1)
*/
package main

import (
	"fmt"
)

func main() {

	str := "abcdcba"

	isPalindrome := true
	j := len(str) - 1
	for i := 0; i < len(str); i++ {
		if string(str[i]) != string(str[j]) {
			isPalindrome = false
			break
		}
		j = j - 1
	}
	fmt.Println("is this a palindrome: ", isPalindrome)
}
