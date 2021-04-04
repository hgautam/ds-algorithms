/*
Returns index of first non-repeating char of a string
input string "assscdddffffrg" will return index 0 corresponding to a
*/
package main

import (
	"fmt"
	"strings"
)

func main() {

	str := "assscdddffffrg"

	myMap := make(map[string]int)
	index := len(str) - 1

	//add chars in map

	for i := 0; i < len(str); i++ {
		key := string(str[i])

		if val, ok := myMap[key]; ok {
			myMap[key] = val + 1
		} else {
			myMap[key] = 1
		}
	}

	fmt.Println(myMap)

	for j := 0; j < len(str); j++ {
		newKey := string(str[j])
		val := myMap[newKey]
		//get the index of val 1
		if val == 1 {
			idx := strings.Index(str, newKey)
			if idx == 0 {
				index = 0
				break
			}
			if index > idx {
				index = idx
			}
		}
	}
	fmt.Println(index)
}
