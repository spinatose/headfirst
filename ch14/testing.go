package main

import (
	"fmt"
	"ch14.com/testing/prose"
)

func main () {
	fruits := []string { "apples", "bananas", "fresas", "ceresas"}
	fmt.Println(prose.JoinWithCommas(fruits))
}
