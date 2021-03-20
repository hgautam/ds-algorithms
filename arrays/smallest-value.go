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
	sort.Ints(a)

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

	fmt.Println(a)
	//fmt.Println(sum)
	//fmt.Println("The smallest value that can be returned: ", smallestVal)
	//fmt.Println("The largest value on the array: ", largestVal)
	fmt.Println("The smallest that cant be returned: ", nonReturnable)
}
