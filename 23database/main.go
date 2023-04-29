package main

import (
	"database/router"
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("MongoDb API")
	r := router.Router()
	fmt.Println("Server is getting started ....")

	log.Fatal(http.ListenAndServe(":4000", r))
	fmt.Println("Listen on port 4000")
}
