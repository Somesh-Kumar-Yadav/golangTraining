package main

import "fmt"

func main() {
	fmt.Println("Struct in golang")
	// No inheritanc in golang, no super or parent

	somesh := User{"Somesh", "someshkumar@gmail.com", true, 24}
	fmt.Println(somesh)                                               // {Somesh someshkumar@gmail.com true 24}
	fmt.Printf("Detials %+v\n", somesh)                               // Detials {Name:Somesh Email:someshkumar@gmail.com Status:true Age:24}
	fmt.Printf("name is %v email is %v\n", somesh.Name, somesh.Email) // name is Somesh email is someshkumar@gmail.com

}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

// Struct needs to be exported out of first letter is Capital
