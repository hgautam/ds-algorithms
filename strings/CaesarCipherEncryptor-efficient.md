/*
Returns a string based on a key
e.g. given key = 2, "xyz" input will result in "zab" output
time complexity: o(n)
space complexity: o(n)
*/
package main

import (
	"fmt"
	"strings"
)

func main() {

	//var sb strings.Builder
	str := "az"
	key := 2
	runes := []rune(str)
	fmt.Println("Input string is: ", str)
	alphabet := "abcdefghijklmnopqrstuvwxyz"

	for i, char := range runes {
		fmt.Println("i is: ", i)
		index := strings.Index(alphabet, string(char))
		fmt.Println("index is: ", index)

		if index == -1 {
			fmt.Println("alphabet not found")
			break
		}

		newIndex := (index + key) % 26

		runes[i] = rune(alphabet[newIndex])
	}

	fmt.Println(string(runes))

}
