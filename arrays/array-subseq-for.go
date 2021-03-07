/*
Use of a for loop based solution to check if a given array is subsequence of a bigger array
*/
package main

import (
	"fmt"
)

func main() {

	a := [8]int{5, 1, 22, 25, 6, -1, 8, 10}
	s := [4]int{1, 6, -1, 10}

	sIdx := 0

	for i := 0; i < len(a); i++ {
		if sIdx == len(s) {
			break
		}

		if s[sIdx] == a[i] {
			sIdx += 1
		}
	}
	fmt.Println(sIdx == len(s))
	//return sIdx == len(s)
}
