package main

import (
	"log"
	"net/http"
)

func main() {
}

func init() {
	println("INIT MAIN")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
