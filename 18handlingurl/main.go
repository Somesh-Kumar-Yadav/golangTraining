package main

import (
	"fmt"
	"net/url"
)

const myurl string = "https://lco.dev:3000/learn?coursename=reactjs&paymentid=gibberishvalue"

func main() {
	fmt.Println("Url handling")

	fmt.Println(myurl)

	// parsing

	result, err := url.Parse(myurl)
	if err != nil {
		panic(err)
	}

	fmt.Println(result.Scheme)
	fmt.Println(result.Host)
	fmt.Println(result.Path)
	fmt.Println(result.Port())
	fmt.Println(result.RawQuery)

	qparams := result.Query()
	fmt.Printf("Type is %T\n", qparams)
	fmt.Println(qparams["coursename"])

	partsOfUrl := &url.URL{ // do not forgot &
		Scheme:   "https",
		Host:     "lco.dev",
		Path:     "/learn",
		RawQuery: "user=somesh",
	}

	anotherUrl := partsOfUrl.String()
	fmt.Println(anotherUrl)

}
