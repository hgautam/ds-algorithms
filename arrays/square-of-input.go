/*
Takes a sorted array as input and returns a sorted array of square root of each integer in the array
Time complexity of o(n)
Space complexity of o(1)
*/
package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Hello, playground")

	a := []int{-5, -4, -3, -2, -1}

	var sqauaredArr []int

	for i := 0; i < len(a); i++ {
		j := a[i] * a[i]
		sqauaredArr = append(sqauaredArr, j)
	}
// this will help we are dealing with an array of negative numbers
	sort.Ints(sqauaredArr)
	fmt.Println(sqauaredArr)
}
