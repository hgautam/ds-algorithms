/*
return the smallest piece of change that cant be returned
Given the array values: 5, 7, 1, 1, 2, 3, 22
The output will be: 20
The smallest value that cannot be returned based on the combination of array values
*/
package main

import (
	"fmt"
	"sort"
)

func main() {

	a := []int{5, 7, 1, 1, 2, 3, 22}
	
	// if emptry array is passed, return 1
	if len(a) < 1 {
		return 1
	}
	// if only one value passed in array, return 1 if the value is greater than 1
	if len(a) == 1 {
		if a[0] > 1 {
			return 1
		}
	}
	sort.Ints(a)
	fmt.Println(a)

	smallestVal := a[0]
	nonReturnable := smallestVal + 1
	//largestVal := len(a)

	sum := 0
	for i := 0; i < len(a); i++ {
		sum += a[i]
		if a[i] == nonReturnable || a[i] < nonReturnable {
			//sum += a[i]
			fmt.Println("a[i]: ", a[i])
			fmt.Println("sum: ", sum)
			nonReturnable = sum + 1
			fmt.Println("non-returnable ", nonReturnable)
		}

	}
	fmt.Println("The smallest that cant be returned: ", nonReturnable)
}
