package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url string = "https://lco.dev"

func main() {
	fmt.Println("Web Request")
	response, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Response is of type: %T\n", response)
	fmt.Println(response)
	defer response.Body.Close()

	dataBytes, err := ioutil.ReadAll(response.Body)

	if err != nil {
		panic(err)
	}
	fmt.Println(string(dataBytes))
}
