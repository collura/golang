package controller

import (
	"encoding/json"
	"github/collura/_model"
	"io"
	"net/http"
)

func PessoaControllerInit() {
	http.HandleFunc("/", handler)
}

func handler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	pessoa := _model.Pessoa{Id: id, Nome: "Betto", Endereco: "Rua Franscisco de Magalhães, 306"}
	json, _ := json.Marshal(pessoa)
	io.WriteString(w, string(json))
}
