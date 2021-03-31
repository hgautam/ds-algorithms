/*
Take a set of strings as input
Print number and string
Input string "AAABBCCDDDD" will result in output 3A,2B,2C,4D
*/
package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var list []string
	str := "AAAAAAAAAAAABBCCDDDDDDDDDD"
	length := 1

	for i := 1; i < len(str); i++ {
		//fmt.Println(i)
		if str[i] == str[i-1] {
			length += 1
		} else {
			//add the string to the list
			if length > 9 {
				list = append(list, strconv.Itoa(9), string(str[i-1]))
				//mod of 9
				modLen := length % 9
				list = append(list, strconv.Itoa(modLen), string(str[i-1]))
				length = 1
			} else {
				list = append(list, strconv.Itoa(length), string(str[i-1]))
				//reset the length back to 1 for the future run
				length = 1
			}
		}
		//handle the last element of the input
		if i == len(str)-1 {
			if length > 9 {
				list = append(list, strconv.Itoa(9), string(str[i-1]))
				//mod of 9
				modLen := length % 9
				list = append(list, strconv.Itoa(modLen), string(str[i-1]))
			} else {
				list = append(list, strconv.Itoa(length), string(str[i-1]))
			}
		}
	}
	//fmt.Println(list)
	result := strings.Join(list, "")
	fmt.Println(result)

}
