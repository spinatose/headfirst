package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	defer reportPanic()
	scanDirectory("my_directory")
}

func reportPanic() {
	p := recover()
	if p == nil {
		return 
	}

	err, ok := p.(error)
	if ok {
		fmt.Println(err)
	} else {
		panic(p)
	}
}

func scanDirectory(path string) {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		panic(err)
	}
	
	for _, file := range files {
		filepath := path + "/" + file.Name()
		if file.IsDir() {
			fmt.Println(filepath)
			scanDirectory(filepath)
		} else {
			fmt.Println(filepath)
		}
	}
}