package main

import (
	"fmt"
	"log"

	"ch10.com/encaps/dates"
)

func main () {
	event := dates.Event { Title: "Momma's bday" }
	event.SetMonth(1)
	event.SetDay  (1)
	event.SetYear (2001)

	date := dates.Date { }
	fmt.Println(date)

	err := date.SetYear(2021) // date automatically gets passed as pointer in setter
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(date)

	err = date.SetYear(1900) // date automatically gets passed as pointer in setter
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(date)

	err = date.SetMonth(5) // date automatically gets passed as pointer in setter
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(date)

	err = date.SetDay(30) // date automatically gets passed as pointer in setter
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(date)
	fmt.Println(event)
}
