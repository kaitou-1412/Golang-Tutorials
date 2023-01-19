package main

import "fmt"

func main() {
	fmt.Println("Welcome to loops in golang")
	days := []string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday"}
	fmt.Println(days)
	fmt.Println("----------Type 1-------------")
	//Type 1
	for d := 0; d < len(days); d++ {
		fmt.Println(days[d])
	}
	fmt.Println("----------Type 2-------------")
	//Type 2
	for i := range days {
		fmt.Println(days[i])
	}
	fmt.Println("-----------Type 3------------")
	//Type 3
	for index, day := range days {
		fmt.Printf("index is %v and day is %v\n", index, day)
	}
	fmt.Println("------------comma ok-----------")
	// comma ok
	for _, day := range days {
		fmt.Printf("index is _ and day is %v\n", day)
	}
	fmt.Println("-----------Type 4------------")
	//Type 4
	rogueValue := 1
	for rogueValue <= 10 {
		fmt.Println("Value is", rogueValue)
		rogueValue++
	}
	fmt.Println("-----------break------------")
	// break
	rogueValue = 1
	for rogueValue <= 10 {
		if rogueValue == 5 {
			break
		}
		fmt.Println("Value is", rogueValue)
		rogueValue++
	}
	fmt.Println("------------continue-----------")
	// continue
	rogueValue = 1
	for rogueValue <= 10 {
		if rogueValue == 5 {
			rogueValue++
			continue
		}
		fmt.Println("Value is", rogueValue)
		rogueValue++
	}
	fmt.Println("------------goto-----------")
	// goto
	rogueValue = 1
	for rogueValue <= 10 {
		if rogueValue == 5 {
			goto ruban
		}
		fmt.Println("Value is", rogueValue)
		rogueValue++
	}
ruban:
	fmt.Println("Jumped on to Ruban")
}
