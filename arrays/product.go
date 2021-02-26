/*
find the product of all the numbers in an array
*/
package main

import (
	"fmt"
)

func main() {

	var array = [5]int{1, 2, 3, 4, 5}
	product := 1

	for i := 0; i < len(array); i++ {
		product = product * array[i]
	}

	fmt.Println("Product is: ", product)
}
