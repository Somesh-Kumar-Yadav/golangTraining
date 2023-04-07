package main

import "fmt"

func main() {
	fmt.Println("Welcome to array in golang")
	var fruitList [4]string
	fruitList[0] = "Array"
	fruitList[1] = "Tomato"
	fruitList[3] = "Peach"
	fmt.Println("FruitList is: ", fruitList)

	var vegList = [3]string{"potato", "beans", "mushroom"}

	fmt.Println(len(vegList))

}
