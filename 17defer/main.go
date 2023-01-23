package main

import "fmt"

func main() {
	// follows LIFO order in case of multiple defers
	// one of the use cases of defer is for closing files
	defer fmt.Println("Hello world")
	defer fmt.Println("One")
	defer fmt.Println("\nTwo")
	fmt.Println("Welcome to defer in golang")
	myDefer()
}

func myDefer() {
	for i := 0; i < 5; i++ {
		defer fmt.Print(i)
	}
}
