/*
o (n+ m) time and c spcae complexity. n is number of chars and m is the length of document
c is the number of unique chars
*/
package main

import (
	"fmt"
)

func main() {
	characters := "aheaolabbhb"
	document := "hello"
	flag := true

	characterCounts := map[rune]int{}

	for _, char := range characters {
		characterCounts[char] = characterCounts[char] + 1
	}

	fmt.Println(characterCounts)

	//see if the document string is found in the char set

	for _, char := range document {
		if characterCounts[char] == 0 {
			flag = false
		}

		characterCounts[char] = characterCounts[char] - 1
	}
	fmt.Println(flag)
}
