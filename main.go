package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	fmt.Println("Hello World")

	// Create a file
	file, err := os.Create("test.txt")
	if err != nil {
		log.Fatal(err)
	}
	file.WriteString("Hello World")
	file.Close()

	// Read a file
	stream, err := ioutil.ReadFile("test.txt")
	if err != nil {
		log.Fatal(err)
	}
	readString := string(stream)
	fmt.Println(readString)
}
