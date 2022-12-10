package controller

import (
	"net/http"

	"github.com/golang-jwt/jwt"
)

func PessoaControllerInit() {
	http.HandleFunc("/", handler)
}

func handler(w http.ResponseWriter, r *http.Request) {
	//	id := r.URL.Query().Get("id")
	tokenStr := r.Header.Get("Authorization")
	print(tokenStr)
	token, _ := jwt.Parse(tokenStr, nil)
	if token == nil {
		return
	}
	claims, _ := token.Claims.(jwt.MapClaims)
	print(claims)
}
