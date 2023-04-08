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

// go list -m all (all dependencies in this module)

// go list -m -versions github.com/gorilla/mux

// go mod tidy (npm i)

// go mod why github.com/gorilla/mux

// go mod graph

// go mod edit -go 1.16

// go mod edit -module 1.16

// go mod vendor (give folder vendor like node_module)

// go run -mod=vendor main.go
