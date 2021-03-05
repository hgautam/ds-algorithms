/*
Use a sorted array to get an efficient solution for space complexity of o(1)
Time complexity in this case will be o(n log n)
*/
package main

import (
	"fmt"
	"sort"
)

func main() {

	var a = []int{3, 5, -4, 8, 11, 1, -1, 6}
	total := 10
	var array []int
	//printed unsorted array
	fmt.Println(a)
	sort.Ints(a)
	//print sorted array
	fmt.Println(a)

	left := 0
	right := len(a) - 1

	fmt.Println("left most value is: ", left)
	fmt.Println("right most value is: ", right)

	for left < right {
		currentSum := a[left] + a[right]
		if currentSum == total {
			array = []int{a[left], a[right]}
			//return []int{a[left], a[right]}
			// the return statement will break this loop
			break
			//fmt.Println(array{a[left], a[right]})
		} else if currentSum < total { //move to the left pointer if sum is less than required total
			left += 1
		} else { // move the right cursor if the sum is more than the required total
			right -= 1
		}
	}
	fmt.Println(array)
  //this will return an empty array to the caller
  //return []int
}
