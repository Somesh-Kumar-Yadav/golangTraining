package main

import "fmt"

func main() {
	fmt.Println("Control flow")

	loginCount := 11

	var result string
	if loginCount < 10 {
		result = "Regular User"
	} else if loginCount > 18 {
		result = "Nothing"
	} else {
		result = "10"
	}

	if num := 11; num < 10 {
		fmt.Println("111")
	} else {
		fmt.Println("333")
	}

	fmt.Println(result)

}
