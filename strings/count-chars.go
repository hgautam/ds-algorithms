/*
Take a set of strings as input
Print number and string
Input string "AAABBCCDDDD" will result in output 3A,2B,2C,4D
*/
package main

import (
	"fmt"
	//"strings"
)

func main() {
	str := "AAABBCCDDDD"

	myMap := make(map[string]int)

	for i := 0; i < len(str); i++ {
		//fmt.Println(string(str[i]))
		key := string(str[i])

		//find the key
		//if it is there increment it
		if val, ok := myMap[key]; ok {
			//increment the value
			myMap[key] = val + 1
		} else {
			myMap[key] = 1
		}
	}
	fmt.Println(myMap)
}
