package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Welcome to slice")

	var fruitList = []string{"Apple", "Tomato", "Peach"} //need to initialize also in this syntax
	fmt.Printf("Type of fruitList is: %T\n", fruitList)

	fruitList = append(fruitList, "Mango", "Banana")

	fmt.Println(fruitList)
	fruitList = append(fruitList[1:])
	fmt.Println(fruitList)

	highScore := make([]int, 4)
	highScore[0] = 234
	highScore[1] = 344
	highScore[2] = 335
	highScore[3] = 265

	highScore = append(highScore, 555, 867, 443)

	fmt.Println(highScore)

	fmt.Println(sort.IntsAreSorted(highScore))
	sort.Ints(highScore)

	fmt.Println(highScore)

	fmt.Println(sort.IntsAreSorted(highScore))

}
