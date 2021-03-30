package main

import (
	"log"
	"net/http"
)

func main () {
	http.HandleFunc("/hello", viewHandler)
	err := http.ListenAndServe("localhost:8080", nil)
	log.Fatal(err)
}

func viewHandler(writer http.ResponseWriter, request *http.Request) {
	message := []byte("Holo! de INternet")
	_, err := writer.Write(message)
	if err != nil {
		log.Fatal(err)
	}
}

