package main

import (
	"./datafile"
	"fmt"
	"log"
	"sort"
)

func main() {
	lines, err := datafile.GetStrings("./votes.txt")

	if err != nil {
		log.Fatal(err)
	}

	votes := make(map[string]int)
	for _, line := range lines {
		votes[line]++ // adds the line if not there and increments the count to 1 on init if missing, otherwise just increments for already present name
	}

	// for key, value := range votes {
	// 	fmt.Printf("%s: %v\n", key, value)
	// }

	// we want to sort alphabetically by name
	var names []string

	for key := range votes {
		names = append(names, key)
	}

	sort.Strings(names)

	for _, name := range names {
		fmt.Printf("%s: %v\n", name, votes[name])
	}
}
