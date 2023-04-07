package main

import "fmt"

func main() {
	fmt.Println("Maps in golang")

	languages := make(map[string]string)

	languages["JavaScript"] = "JS"
	languages["Ruby"] = "RB"
	languages["Python"] = "PY"

	fmt.Println(languages)
	fmt.Println(languages["JavaScript"])
	delete(languages, "Ruby")

	fmt.Println(languages)

	// loops are interesting in golang

	for _, value := range languages {
		fmt.Println(value)
		fmt.Printf("key v, value %v\n", value)
	}
}
