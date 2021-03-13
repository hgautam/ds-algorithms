package main

import (
	"fmt"
)

func main() {
	teams := map[string]int{
		"HTML":   3,
		"PYTHON": 6,
		"GO":     21,
		"JAVA":   18,
	}

	//iterate over the keys and values
	for key, value := range teams {
		fmt.Println(key, value)
	}

	//search for a key and update it
	if val, ok := teams["GO"]; ok {
		//do something here
		fmt.Println("found", val)
		teams["GO"] = 27
	}

	//iterate over the keys and values
	for key, value := range teams {
		fmt.Println(key, value)
	}
}
