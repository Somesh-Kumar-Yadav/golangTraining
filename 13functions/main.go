package main

import "fmt"

func main() {
	fmt.Println("Welcome to functions")
	greeter()
	greeterTwo()
	result := add(1, 2)
	resultall, message := addAll(1, 2, 3, 4, 5)
	fmt.Println(result)
	fmt.Println(message)
	fmt.Println(resultall)

}

func greeterTwo() {
	fmt.Println("Another method")

}

func add(a int, b int) int {
	// fmt.Println(a + b)
	result := a + b
	return result
}

func addAll(values ...int) (int, string) {
	total := 0
	for _, value := range values {
		total += value
	}
	return total, "Somesh"
}

func greeter() {
	fmt.Println("Namastey Golang")
}
