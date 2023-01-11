package main

import "fmt"

func main() {
  fmt.Println("Welcome to a class on pointers")

  // var ptr *int
  // fmt.Println("Value of ptr is ", ptr)

  myNumber := 23
  var ptr = &myNumber
  fmt.Println("Reference of myNumber = Value of ptr = ", ptr)
  fmt.Println("Value of myNumber = Value pointed by ptr = ", *ptr)
}
