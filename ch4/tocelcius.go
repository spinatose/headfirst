package main

import (
	"./input"
	"fmt"
	"log"
)

func main() {
	fmt.Print("Enter a temp in Fahrenheit: ")
	fahr, err := input.GetFloat()
	if err != nil {
		log.Fatal(err)
	}
	celcius := (fahr - 32) * 5 / 9
	fmt.Printf("%0.2f degrees Celcius\n", celcius)
}
