package main

import "fmt"

func main () {
	next_int := int_seq()
	fmt.Println(next_int()) // now next_int has like it's own shadow object that internalizes the state of i
	fmt.Println(next_int())
	prev_int := next_int
	next_int = int_seq() // this is a different closure returned, so we should see another init message along with a print out of 1 again
	fmt.Println(next_int())
	// however, now prev_int refers to the original closure, we shall see 3 print out with this call to it
	fmt.Println(prev_int())
	fmt.Println(byte(0x0b))
}

func int_seq() func() int {
	i := 0 
	fmt.Println("Initial call to int_seq which returns a closure- what is i? ", i)
	return func () int {
		i++
		return i 
	}
}