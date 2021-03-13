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
	/* an array with 2 rows and 2 columns*/
	var a = [3][2]string{{"HTML", "C#"}, {"C#", "PYTHON"}, {"PYTHON", "HTML"}}
	var result = [3]int{0, 0, 1}

	//myMap := map[string]int{}

	for i := 0; i < len(a); i++ {
		var str []string
		for j := 0; j < len(a[0]); j++ {
			//fmt.Println(a[i][j])
			str = append(str, a[i][j])
		}
		//fmt.Println("home team: ", str[0])
		//fmt.Println("away team: ", str[1])
		if result[i] == 1 {
			fmt.Println("winner is: ", str[0])
			if key, ok := myMap[str[0]]; ok {
				//update the value
				key[]
			}
		} else {
			fmt.Println("winner is: ", str[1])
		}
		//if key, ok := myMap["foo"]; ok {
		//do something here
		//}
		fmt.Println("Record set complete")
	}

}
