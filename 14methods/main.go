package main

import "fmt"

func main() {
	fmt.Println("Welcome to method")
	somesh := User{"Somesh", "someshkumar@gmail.com", true, 24}
	fmt.Println(somesh)
	somesh.GetStatus()                                                // {Somesh someshkumar@gmail.com true 24}
	fmt.Printf("Detials %+v\n", somesh)                               // Detials {Name:Somesh Email:someshkumar@gmail.com Status:true Age:24}
	somesh.NewMail()
	fmt.Printf("name is %v email is %v\n", somesh.Name, somesh.Email) // name is Somesh email is someshkumar@gmail.com
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
	// oneAge int {} not exportable }
}

func (u User) GetStatus() {
	fmt.Println(u.Status)
}

func (u User) NewMail() {
	u.Email = "yadav.somesh@gmail.com"
	fmt.Println(u.Email)
}
