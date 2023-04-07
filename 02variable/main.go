package main

import "fmt"

const LastName string = "Yadav"

func main() {
	var name string = "Somesh"
	var age int16 = 16
	var height float32 = 5.45
	var weight uint32 = 60

	fmt.Println(name)
	fmt.Printf("%T \n", name)

	fmt.Println(age)
	fmt.Printf("%T \n", age)

	fmt.Println(height)
	fmt.Printf("%T \n", height)

	fmt.Println(weight)
	fmt.Printf("%T \n", weight)

	loveInterest := "Namrta Yadav"

	fmt.Println(loveInterest)
	fmt.Printf("%T \n", loveInterest)


}
