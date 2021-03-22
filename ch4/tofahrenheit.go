package main

import (
	"./input"
	"fmt"
	"log"
)

func main() {
	fmt.Print("Enter a temp in Celcius: ")
	celcius, err := input.GetFloat()
	if err != nil {
		log.Fatal(err)
	}
	fahr := (celcius * 9 / 5) + 32
	fmt.Printf("%0.2f degrees Fahrenheit\n", fahr)
}
