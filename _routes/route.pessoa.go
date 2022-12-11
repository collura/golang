package _routes

import (
	"github/collura/controller"
	"net/http"
)

func PessoaControllerInit() {
	http.HandleFunc("/pessoa", controller.HandlerPessoa)
}
