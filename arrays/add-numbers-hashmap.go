/*
Uses a hash map to find to array integers that add to a sum
This is o(n) time complexity solution
 */
package main

import (
	"fmt"
)

func main() {

	//create empty map
	var myMap = map[int]int{}

	var arr []int

	var a = []int{3, 5, -4, 8, 11, 1, -1, 6}
	total := 10
	number := 0

	//iterate through the array to see if we get the number that makes the sum
	for i := 0; i < len(a); i++ {
		number = total - a[i]
		if _, ok := myMap[number]; ok {
			//myMap[number] = number
			arr = []int{a[i], number}

		} else {
			fmt.Println("adding number", a[i])
			myMap[a[i]] = a[i]
		}
	}

	fmt.Println(arr)
	fmt.Println(myMap)
}
