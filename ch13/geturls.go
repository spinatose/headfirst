package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func main() {
	go responseSize("http://www.rt.com")
	go responseSize("http://www.rantingly.com")
	go responseSize("http://www.bing.com")

	time.Sleep(5 * time.Second)
	fmt.Println("Se acabó...")
}

func responseSize(url string) {
	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Getting", url, ": ", len(body))
}
