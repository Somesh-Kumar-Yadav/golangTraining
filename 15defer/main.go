package main

import (
	"fmt"
)

func main() {
	fmt.Println("Defer")
	defer fmt.Println("world")
	defer fmt.Println("world2")
	fmt.Println("Hello")
	myDefer()
}

func myDefer() {
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}
}
