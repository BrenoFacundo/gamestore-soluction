package model 

type Game struct {
	ID int `json:"id"`
	Name string `json:"nome"`
	Price float32 `json:"preco"`
	Platform string `json:"plataforma"`
	Description string `json:"descricao"`
}

//PREGUIÇA DE TROCAR OS CAMPOS NO DB PRA INGLES