/*
Given an input string "aaabbb" check if we can produce "ab"
*/

package main

import (
	"fmt"
)

func main() {
	str := "aheaolabbhb"
	outputStr := "hello"
	flag := false

	// if output is zero length string then return true
	if len(outputStr) < 1 {
		//return true
		flag = true
		return
	}

	myMap1 := make(map[string]int)
	myMap2 := make(map[string]int)
	createaMap(str, myMap1)
	createaMap(outputStr, myMap2)
	fmt.Println(myMap1)
	fmt.Println(myMap2)

	//itereate over output string and find if input string has those chars
	for i := 0; i < len(outputStr); i++ {
		key := string(outputStr[i])
		//get the number of time key exists in the map
		val := myMap2[key]
		val1 := myMap1[key]

		//if val1 is >= val we are goog
		if val1 >= val {
			fmt.Println("life is good")
			flag = true
		} else {
			fmt.Println("life sucks...")
			flag = false
			break
		}
	}
	fmt.Println("flag is: ", flag)
}

func createaMap(str string, myMap map[string]int) {
	for i := 0; i < len(str); i++ {
		key := string(str[i])
		//fmt.Println(key)
		if val, ok := myMap[key]; ok {
			//increment the value
			myMap[key] = val + 1
		} else {
			myMap[key] = 1
		}
	}
}
