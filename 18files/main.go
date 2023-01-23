package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("Welcome to files in golang")
	filename := "./myFile.txt"
	content := "This needs to go in a file."
	file, err := os.Create(filename)
	checkNilError(err)
	defer file.Close()
	length, err := io.WriteString(file, content)
	checkNilError(err)
	fmt.Println("length is:", length)
	readFile(filename)
}

func readFile(filename string) {
	dataBytes, err := ioutil.ReadFile(filename)
	checkNilError(err)
	data := string(dataBytes)
	fmt.Println("Text data inside the file is:\n", data)
}

func checkNilError(err error) {
	if err != nil {
		//shuts down the execution of the program and shows the error
		panic(err)
	}
}
