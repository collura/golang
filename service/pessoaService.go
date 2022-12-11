package service

import (
	"encoding/json"
	"github/collura/_model"
)

func Handler(id int64) string {
	pessoa := _model.Pessoa{Id: id + 1, Nome: "betto", Endereco: "Rua simoes delgado, 306"}
	str, _ := json.Marshal(pessoa)
	return string(str)
}
