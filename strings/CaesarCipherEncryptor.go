/*
Returns a string based on a key
e.g. given key = 2, "xyz" input will result in "zab" output
Time efficieny should be o(n)
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
		"x": 24,
		"y": 25,
		"z": 26,
	}
	var sb strings.Builder
	str := "xyz"
	myKey := 2
	fmt.Println("Input string is: ", str)
	for i := 0; i < len(str); i++ {
		charVal := string(str[i])
		//find the value for a given string
		value := alphaMap[charVal]
		//fmt.Println("value for string is: ", value)

		myRealKey := value + myKey

		if myRealKey > 26 {
			myRealKey = myRealKey % 26
		}

		//find a value in the map and print its key

		for key, val := range alphaMap {
			//fmt.Println("Key:", key, "=>", "value:", val)
			if val == myRealKey {
				//fmt.Println("found")
				//fmt.Println("Return string is: ", key)
				sb.WriteString(key)
				break
			}
		}
	}
	fmt.Println("Return string is: ", sb.String())
}
