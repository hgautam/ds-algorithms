/*
A multi-dimensional array represnts teams in a tournament
second array denotes the result for the home team
We determine the winner of the contest
Time complexity: o(n^2) 
space complexity:o(k). n is number of competitions, k is number of teams
*/
package main

import (
	"fmt"
)

func main() {
	/* an 2 D array with 3 rows and 2 columns*/
	var competitions = [3][2]string{{"HTML", "C#"}, {"C#", "PYTHON"}, {"PYTHON", "HTML"}}
	var results = [3]int{0, 0, 1}

	const HOME_TEAM_WON = 1
	bestTeam := " "
	scores := map[string]int{bestTeam: 0}

	for i := 0; i < len(competitions); i++ {
		var str []string
		for j := 0; j < len(competitions[0]); j++ {
			//fmt.Println(competitions[i][j])
			str = append(str, competitions[i][j])
		}
		winningTeam := str[1]
		if results[i] == HOME_TEAM_WON {
			fmt.Println("winner is home team: ", str[0])
			winningTeam = str[0]

		}
		calScore(winningTeam, scores)
		// set winning team as the best team
		if scores[winningTeam] > scores[bestTeam] {
			bestTeam = winningTeam
		}
	}
	fmt.Println("scores has these entries: ", scores)
	fmt.Println("final winner is: ", bestTeam)
}

func calScore(winner string, scores map[string]int) {
	scores[winner] += 3
}
