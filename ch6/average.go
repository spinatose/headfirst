package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	var args []string = os.Args[1:]
	var nums []float64 
	sum := 0.0

	for _, arg := range args {
		num, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			log.Fatal(err)
		}

		nums = append(nums, num)
		sum += num
	}

	fmt.Printf("Total: %.2f\n", sum)
	fmt.Printf("Average: %.2f\n", sum/ float64(len(args)))
	fmt.Printf("Average from average func: %.2f\n", average(nums...))
}

func average(numbers ...float64) float64 {
	sum := 0.0

	for _, num := range numbers {
		sum += num
	}

	return sum / float64(len(numbers))
}