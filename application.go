package main

import (
	"github/collura/controller"
	"log"
	"net/http"
)

func main() {
	controller.PessoaControllerInit()
	log.Fatal(http.ListenAndServe(":8080", nil))
}

type Claims struct {
	Username string `json:"username"`
}
