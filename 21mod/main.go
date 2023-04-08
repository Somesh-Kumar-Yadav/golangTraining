package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Hello mod")
	gretter()
	r := mux.NewRouter()
	r.HandleFunc("/", serveHome).Methods("GET")

	log.Fatal(http.ListenAndServe(":4002", r))
}

// go get -u github.com/gorilla/mux

func gretter() {
	fmt.Println("Namastey")
}

func serveHome(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte("<h1>Hello Somesh</h1>"))
}
