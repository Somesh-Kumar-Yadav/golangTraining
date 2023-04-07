package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

const ur string = "http://localhost:3000"

func main() {
	fmt.Println("Welcome to web verb")
	// PerformGetRequest()
	// PerformPostJsonRequest()
	PerformPostEncodedRequest()

}

func PerformGetRequest() {
	response, err := http.Get(ur + "/get")
	if err != nil {
		panic(err)
	}

	defer response.Body.Close()
	fmt.Println("status code: ", response.StatusCode)
	fmt.Println("content length: ", response.ContentLength)

	dataByte, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(dataByte))

	var responseString strings.Builder
	byteCount, _ := responseString.Write(dataByte)
	fmt.Println(byteCount)
	fmt.Println(responseString.String())

}

func PerformPostJsonRequest() {
	requestBody := strings.NewReader(`
		{
			"name":"somesh",
			"price":1
		}
	`)
	response, err := http.Post(ur+"/post", "application/json", requestBody)
	if err != nil {
		panic(err)
	}
	databyte, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(databyte))
	defer response.Body.Close()

}

func PerformPostEncodedRequest() {
	data := url.Values{}

	data.Add("firstname", "somesh")
	data.Add("lastname", "name")

	response, err := http.PostForm(ur+"/postform", data)
	if err != nil {
		panic(err)
	}
	databyte, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(databyte))
	defer response.Body.Close()

}
