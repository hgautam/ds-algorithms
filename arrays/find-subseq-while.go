/*
Use of a while loop based solution to check if a given array is subsequence of a bigger array
*/
package main

import (
	"fmt"
)

func main() {

	a := [8]int{5, 1, 22, 25, 6, -1, 8, 10}
	s := [4]int{1, 6, -1, 10}

	arrIdx := 0
	sIdx := 0

	for arrIdx < len(a) && sIdx < len(s) {
		if s[sIdx] == a[arrIdx] {
			sIdx += 1
		}
		arrIdx += 1
	}

	fmt.Println(sIdx == len(s))
	//return sIdx == len(s)
}
