/*
Find two numbers that match a target sum
Sum cannot be arrived at by add two different integer twice
You cant arrive at the solution by adding 5 to itself
*/
package main

import (
	"fmt"
)

func main() {
	var a = []int{3, 5, -4, 8, 11, 1, -1, 6}

	sum := 10

	fmt.Println("length of the array", len(a))
out:
	for i := 0; i < len(a); i++ {
		for j := i; j < len(a); j++ {
			// skip adding the integer to itself
			if i == j {
				continue
			}
			if a[i]+a[j] == sum {
				fmt.Println(a[i])
				fmt.Println(a[j])
				break out
			}
		}
	}
}
