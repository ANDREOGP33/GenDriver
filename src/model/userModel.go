package model

type User struct {
	Clinica string
	Cnpj    string
	Senha   string
}

type UserAutenticado struct {
	Cnpj      string
	SenhaHash string
}


