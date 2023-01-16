package main

import "fmt"

func main() {
	fmt.Println("Structs in golang")
	// no inheritance in golang; No super or parent
	ruban := User{"ruban", "ruban@go.dev", true, 23}
	fmt.Println(ruban)
	fmt.Printf("Ruban details are: %+v\n", ruban)
	fmt.Printf("Name is %v and Email is %v\n", ruban.Name, ruban.Email)
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
