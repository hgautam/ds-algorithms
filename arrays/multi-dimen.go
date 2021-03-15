/*
A multi-dimensional array represnts teams in a tournament
second array denotes the result for the home team
We determine the winner of the contest
*/
package main

import (
	"fmt"
)

func main() {
	/* an 2 D array with 3 rows and 2 columns*/
	var a = [3][2]string{{"HTML", "C#"}, {"C#", "PYTHON"}, {"PYTHON", "HTML"}}
	var result = [3]int{0, 0, 1}

	scores := map[string]int{}
	var winner string

	for i := 0; i < len(a); i++ {
		var str []string
		for j := 0; j < len(a[0]); j++ {
			//fmt.Println(a[i][j])
			str = append(str, a[i][j])
		}
		if result[i] == 1 {
			fmt.Println("winner is home team: ", str[0])
			if val, ok := scores[str[0]]; ok {
				//update the value
				scores[str[0]] = val + 3
			} else {
				scores[str[0]] = 3
				winner = str[0]
			}
		} else {
			fmt.Println("winner is away team: ", str[1])
			if val, ok := scores[str[1]]; ok {
				//update the value
				scores[str[0]] = val + 3
				//fmt.Println("value is: ", val)
			} else {
				scores[str[1]] = 3
				winner = str[1]
			}
		}
		//fmt.Println("Record set complete")
	}
	fmt.Println("scores has these entries: ", scores)
	fmt.Println("final winner is: ", winner)
}
