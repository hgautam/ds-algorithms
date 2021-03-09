/*
Sorted array input and output is sorted squares
Non-brute force algorithm
Time complexity o(n)
space complexity o(n)
On time complexity is better than o(n log n)
*/
package main

import (
	"fmt"
)

func main() {

	array := [7]int{-3, -1, 0, 2, 4, 6, 9}
	sortedArray := make([]int, len(array))
	smallerIdx := 0
	largerIdx := len(array) - 1
	// reverse traverse for a sorted array
	// starting with the largest element
	for i := len(array) - 1; i >= 0; i-- {
		//fmt.Println(abs(array[i]))
		smallerVal := array[smallerIdx]
		largerVal := array[largerIdx]
		//start the comparison here
		if abs(smallerVal) > abs(largerVal) {
			sortedArray[i] = smallerVal * smallerVal
			smallerIdx += 1
		} else {
			sortedArray[i] = largerVal * largerVal
			largerIdx -= 1
		}
	}
	fmt.Println(sortedArray)
  //return sortedArray

}

func abs(a int) int {
	if a < 0 {
		return a * -1
    // return -a
    //the above seems to work to
	}
	return a
}
