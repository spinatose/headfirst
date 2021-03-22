package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type page struct {
	url string
	size int
}

func main() {
	respchan := make(chan page)
	urls := []string { "http://www.rt.com", "http://www.rantingly.com", "http://www.bing.com"}

	for _, url := range urls {
		go responseSize(url, respchan)
	}
	
	for range urls {
		resp := <-respchan
		fmt.Printf("Getting %s: %v\n", resp.url, resp.size)
	}

	fmt.Println("Se acabó...")
}

func responseSize(url string, respchan chan page) {
	resp := page { url: url }

	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	resp.size = len(body)
	respchan <- resp
}
