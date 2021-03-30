package main

import (
	"bufio"
	"fmt"
	"log"
	"html/template"
	"net/http"
	"os"
)

type Guestbook struct {
	SignatureCount int
	Signatures []string 
}

func main () {
	http.HandleFunc("/guestbook", viewHandler)
	http.HandleFunc("/guestbook/new", newHandler)
	http.HandleFunc("/guestbook/create", createHandler)

	err := http.ListenAndServe("localhost:8080", nil)
	log.Fatal(err)
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func createHandler(writer http.ResponseWriter, request *http.Request) {
	sig := request.FormValue("signature")
	options := os.O_WRONLY | os.O_APPEND | os.O_CREATE
	file, err := os.OpenFile("signatures.txt", options, os.FileMode(0600))
	check(err)
	_, err = fmt.Fprintln(file, sig)
	check(err)
	err = file.Close()
	check(err)
	http.Redirect(writer, request, "/guestbook", http.StatusFound)
}

func getStrings(fileName string) []string {
	var lines []string
	file, err := os.Open(fileName)
	if os.IsNotExist(err) {
		return nil
	}
	check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	check(scanner.Err())

	return lines
}

func newHandler(writer http.ResponseWriter, request *http.Request) {
	tmplt, err := template.ParseFiles("new.html")
	check(err)
	err = tmplt.Execute(writer, nil)
	check(err)
	err = tmplt.Execute(os.Stdout, nil)
	check(err)
}

func viewHandler(writer http.ResponseWriter, request *http.Request) {
	sigs := getStrings("signatures.txt")

	guestbook := Guestbook {
		SignatureCount: len(sigs),
		Signatures: sigs,
	}

	tmplt, err := template.ParseFiles("view.html")
	check(err)
	err = tmplt.Execute(writer, guestbook)
	check(err)
	err = tmplt.Execute(os.Stdout, guestbook)
	check(err)
}

