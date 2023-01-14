package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Let's learn about slices")

	var fruitList = []string{"Apple", "Tomato", "Peach"}
	fmt.Printf("Type of fruitList is %T\n", fruitList)

	fruitList = append(fruitList, "Mango", "Banana")
	fmt.Println(fruitList)

	fruitList = append(fruitList[1:])
	fmt.Println(fruitList)

	fruitList = append(fruitList[1:3])
	fmt.Println(fruitList)

	highScores := make([]int, 4)
	highScores[0] = 234
	highScores[1] = 345
	highScores[2] = 456
	highScores[3] = 567
	// highScores[4] = 678 -> gives index out range error
	highScores = append(highScores, 555, 666, 321)
	fmt.Println(sort.IntsAreSorted(highScores))
	fmt.Println(highScores)
	sort.Ints(highScores)
	fmt.Println(highScores)
	fmt.Println(sort.IntsAreSorted(highScores))

}
