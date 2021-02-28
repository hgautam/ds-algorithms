/*
Instructor solution for product of two largest numbers in an array
*/
package main

import (
	"fmt"
)

func main() {

	var array = [5]int{10, 3, 2, 7, 12}

	maxIndex1 := -1
	maxIndex2 := -1

	//find the index of the largest number

	for i := 0; i < len(array); i++ {
		if maxIndex1 == -1 || array[i] > array[maxIndex1] {
			maxIndex1 = i
		}
	}

	fmt.Println("index of the largest number is ", maxIndex1)

	// find the index of the second largest number

	for j := 0; j < len(array); j++ {
		if (maxIndex2 == -1) || (array[j] != array[maxIndex1] && array[j] > array[maxIndex2]) {
			maxIndex2 = j
		}
	}

	fmt.Println("Index of the second largest number is ", maxIndex2)

	fmt.Println("The product of two largest numbers is ", array[maxIndex1]*array[maxIndex2])

}
