/*
Take a set of strings as input
Print number and string
Input string "AAABBCCDDDD" will result in output 3A,2B,2C,4D
if the input has a string greater than 9 then it will be split 9 str and remainder str
e.g. 10As will be return 9A1A
*/
package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var list []string
	str := " "
	length := 1

	if len(str) == 1 {
		fmt.Println(strconv.Itoa(length) + str)
		return
	}

	for i := 1; i < len(str); i++ {
		fmt.Println(i)
		if str[i] == str[i-1] {
			length += 1
		} else {
			//add the string to the list
			if length > 9 {
				//call the method
				list = processData(length, string(str[i-1]), list)
				fmt.Println("list is: ", list)
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
				list = processData(length, string(str[i]), list)

			} else {
				list = append(list, strconv.Itoa(length), string(str[i]))
			}
		}
	}
	//fmt.Println(list)
	result := strings.Join(list, "")
	fmt.Println(result)
}

func processData(length int, str string, list []string) []string {
	newLength := length / 9
	if newLength > 0 {
		for j := 1; j < newLength+1; j++ {
			list = append(list, strconv.Itoa(9), str)
		}
	}
	//mod of 9
	// to calculate remainder and add it to the list
	modLen := length % 9
	if modLen > 0 {
		list = append(list, strconv.Itoa(modLen), str)
	}
	return list
}
