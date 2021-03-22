// pass_fail reports whether a grade is passing or failing.
package main

import (
	"./input"
	"fmt"
	"log"
)

func main() {
	fmt.Print("Enter a grade: ")
	grade, err := input.GetFloat()
	if err != nil {
		log.Fatal(err)
	}

	var status string
	if grade >= 70 {
		status = "passing"
	} else {
		status = "failing"
	}

	fmt.Println("A grade of", grade, "is", status)
}
