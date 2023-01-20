package main

import "fmt"

func main() {
	fmt.Println("Welcome to functions in golang")
	greeter()
	result := adder(3, 5)
	fmt.Println("Result is:", result)
	proResult := proAdder(3, 5, 7)
	fmt.Println("ProResult is:", proResult)
	proResult2, proMessage := proAdderWithMessage(3, 5, 7)
	fmt.Println("ProResult2 is:", proResult2, "Message:", proMessage)
}

func adder(valOne int, valTwo int) int {
	return valOne + valTwo
}

func proAdder(values ...int) int {
	total := 0

	for _, val := range values {
		total += val
	}

	return total
}

func proAdderWithMessage(values ...int) (int, string) {
	total := 0

	for _, val := range values {
		total += val
	}

	return total, "Hi there!"
}

func greeter() {
	fmt.Println("Namastey Golang")
}
