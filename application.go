package main

import (
	"github/collura/controller"
	"log"
	"net/http"

	"github.com/golang-jwt/jwt/v4"
)

func main() {
	controller.PessoaControllerInit()
	log.Fatal(http.ListenAndServe(":8080", nil))
}

type Claims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}
