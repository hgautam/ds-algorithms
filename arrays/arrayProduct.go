/*
A brute force algorithm to find the products of two largest elementes of an array
*/
package main

import (
	"fmt"
)

func main() {

	var array = [5]int{1, 2, 4, 5, 6}

	result := 0
	for i := 0; i < len(array); i++ {
		for j := i + 1; j < len(array); j++ {
			if array[i]*array[j] > result {
				result = array[i] * array[j]
			}
		}
	}
	fmt.Println(result)
}
