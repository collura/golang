package _model

type Pessoa struct {
	Id       int64  `json:"id"`
	Nome     string `json:"nome"`
	Endereco string `json:"endereco"`
}
