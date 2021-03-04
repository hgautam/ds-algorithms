/*
Find two numbers that match a target sum
Sum cannot be arrived at by add same integer twice
You cant arrive at the solution by adding 5 to itself
returnArray slice stores the number combination and
can be returned to the caller of this func
the input is {3, 5, -4, 8, 11, 1, -1, 6}
expected return will be {11, -1}
*/
package main

import (
	"fmt"
)

func main() {
	var a = []int{3, 5, -4, 8, 11, 1, -1, 6}

	sum := 10

	var returnArray []int

	//out:
	for i := 0; i < len(a); i++ {
		for j := i + 1; j < len(a); j++ {
			if a[i]+a[j] == sum {
				returnArray = []int{a[i], a[j]}
				// the following snippet creates a slice
				// the slice should be used if you want to store all combinations that lead to the target
				//returnArray = append(returnArray, a[i])
				//returnArray = append(returnArray, a[j])
				//use break when using slice if you want to find first combination that matches the target sum
				//break out
			}
		}
	}
	fmt.Println("return array is: ", returnArray)
}
