/*
Returns a string based on a key
e.g. given key = 2, "xyz" input will result in "zab" output
*/
package main

import (
	"fmt"
	"strings"
)

func main() {

	alphaMap := map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
		"d": 4,
		"e": 5,
		"f": 6,
		"g": 7,
		"h": 8,
		"i": 9,
		"j": 10,
		"k": 11,
		"l": 12,
		"m": 13,
		"n": 14,
		"o": 15,
		"p": 16,
		"q": 17,
		"r": 18,
		"s": 19,
		"t": 20,
		"u": 21,
		"v": 22,
		"w": 23,
		"x": 24,
		"y": 25,
		"z": 26,
	}
	var sb strings.Builder
	str := "xyz"
	key := 5
	fmt.Println("Input string is: ", str)
	for i := 0; i < len(str); i++ {
		charVal := string(str[i])
		//find the value for a given string
		value := alphaMap[charVal]
		//fmt.Println("value for string is: ", value)

		myRealKey := value + key

		if myRealKey > 26 {
			myRealKey = myRealKey % 26
		}

		//find a value in the map and print its key

		for alpahbet, val := range alphaMap {
			//fmt.Println("Alphabet:", alpahbet, "=>", "value:", val)
			if val == myRealKey {
				//fmt.Println("found")
				//fmt.Println("Return string is: ", alpahbet)
				sb.WriteString(alpahbet)
				break
			}
		}
	}
	fmt.Println("Return string is: ", sb.String())
}
