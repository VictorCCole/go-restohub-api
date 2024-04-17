package models

type Pessoa struct {
	ID   int    `json:"id"`
	Nome string `json:"nome"`
}

type Restaurante struct {
	ID   int    `json:"id"`
	Nome string `json:"nome"`
}

type Pedido struct {
	ID            int    `json:"id"`
	PessoaID      int    `json:"pessoa_id"`
	RestauranteID int    `json:"restaurante_id"`
	Descricao     string `json:"descricao"`
}
