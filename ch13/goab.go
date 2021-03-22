package main

import (
	"fmt"
	"time"
)

func main() {
	go printChar('a')
	go printChar('b')

	time.Sleep(time.Second)
	fmt.Println("Finished...")
}

func printChar(letter rune) {
	for i := 0; i < 50; i++ {
		fmt.Printf("%v", string(letter))
	}

	fmt.Println()
}
