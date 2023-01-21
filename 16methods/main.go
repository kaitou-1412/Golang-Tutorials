package main

import "fmt"

func main() {
	fmt.Println("Welcome to methods in golang")
	ruban := User{"ruban", "ruban@go.dev", true, 23}
	fmt.Println(ruban)
	fmt.Printf("Ruban details are: %+v\n", ruban)
	fmt.Printf("Name is %v and Email is %v\n", ruban.Name, ruban.Email)
	ruban.GetStatus()
	ruban.NewMail()
	fmt.Printf("Name is %v and Email is %v\n", ruban.Name, ruban.Email)
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func (u User) GetStatus() {
	fmt.Println("Is user active:", u.Status)
}

func (u User) NewMail() {
	u.Email = "test@go.dev"
	fmt.Println("Wmail of this user is:", u.Email)
}
