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

	// for _, value := range days {
	// 	fmt.Println(value)
	// }

	rougueValue := 1

	for rougueValue < 10 { // similar to while loop
		if rougueValue == 2 {
			goto lco
		}
		fmt.Println(rougueValue)
		rougueValue++
	}

lco:
	fmt.Println("Jumping")

}
