package main

import (
	"fmt"
	"log"
	"./datafile"
)

func main() {
	numbers, err := datafile.GetFloats("data.txt")
	if err != nil {
		log.Fatal(err)
	}

	sum := 0.0
	for _, number := range numbers {
		fmt.Printf("%.2f\n", number)
		sum += number 
	}

	fmt.Printf("Total: %.2f\n", sum)
	fmt.Printf("Average: %.2f\n", sum / float64(len(numbers)))
}

/*
func closeFile (file *os.File) {
	err := file.Close()
	if err != nil {
		log.Fatal(err)
	}
}

func closeScanner (scanner *bufio.Scanner) {
	if scanner.Err() != nil {
		log.Fatal(scanner.Err())
	}
}
*/