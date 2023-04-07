package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Switch")
	rand.Seed(time.Now().UnixNano())
	diceNumber := rand.Intn(6) + 1

	fmt.Println(diceNumber)
	switch diceNumber {
	case 1:
		fmt.Println("Dice value is 1 and you can opwn")
		break
	case 2:
		fmt.Println("you can move 2 spot")
		break
	case 3:
		fmt.Println("you can move 3 spot")
		break
	case 4:
		fmt.Println("you can move 4 spot")
		break
	case 5:
		fmt.Println("you can move 5 spot")
		fallthrough
	case 6:
		fmt.Println("you can move 6 spot")
		break
	default:
		fmt.Println("what is this")
		break
	}

}
