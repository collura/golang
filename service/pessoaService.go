package service

import (
	"encoding/json"
	model "github/collura/_model"
)

func Handler(id string) string {
	pessoa := model.Pessoa{Id: id, Nome: "betto", Endereco: "Rua simoes delgado, 306"}
	str, _ := json.Marshal(pessoa)
	return string(str)
}
