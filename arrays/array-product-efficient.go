/*
find the product of two largest numbers in the array
This algorithm calculates the largest number first
Skips the largest number while calculating the second largest number
Hence making it efficient
*/
package main

import (
	"fmt"
)

func main() {

	var array = [5]int{12, 2, 120, 75, 0}

	maxIndex1 := -1
	maxIndex2 := -1
	number := -1
	number2 := -1
        // find the index of the largest number in the array
	for i := 0; i < len(array); i++ {
		for j := i + 1; j < len(array); j++ {
			if array[i] > array[j] && number < array[i] {
				number = array[i]
				maxIndex1 = i
			}
			if array[j] > array[i] && number < array[j] {
				number = array[j]
				maxIndex1 = j
			}
		}
	}
        //find the index of the second largest number in the array
	for i := 0; i < len(array); i++ {
		if i == maxIndex1 {
			continue
		} else {
			for j := i + 1; j < len(array); j++ {
				if j == maxIndex1 {
					continue
				}
				if array[i] > array[j] && number2 < array[i] {
					number2 = array[i]
					maxIndex2 = i
				}
				if array[j] > array[i] && number2 < array[j] {
					number2 = array[j]
					maxIndex2 = j
				}
			}
		}
	}

	fmt.Println(number)
	fmt.Println(maxIndex1)
	fmt.Println(number2)
	fmt.Println(maxIndex2)
	// multiple the largest and the second largest number to get the product
	product := number * number2
	fmt.Println("the product of two largest numbers is:", product)
}
