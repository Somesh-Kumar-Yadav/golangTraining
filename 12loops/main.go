package main

import "fmt"

func main() {
	fmt.Println("Welcome to loops")
	days := []string{"Sunday", "Tuesday", "Wednesday", "Friday", "Saturday"}

	fmt.Println(days)

	// for d := 0; d < len(days); d++ { // no ++d
	// 	fmt.Println(days[d])
	// }

	// for i := range days {
	// 	fmt.Println(i, days[i])
	// }

	for _, value := range days {
		fmt.Println(value)
	}

}
