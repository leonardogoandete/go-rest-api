package main

import (
	"fmt"
	"github.com/leonardogoandete/go-rest-api/database"
	"github.com/leonardogoandete/go-rest-api/models"
	"github.com/leonardogoandete/go-rest-api/routes"
)

func main() {
	models.Personalidades = []models.Personalidade{
		{Id: 1, Nome: "Albert Einstein", Historia: "Físico teórico, desenvolvedor da teoria da relatividade."},
		{Id: 2, Nome: "Marie Curie", Historia: "Pioneira na pesquisa sobre radioatividade, primeira mulher a ganhar um Prêmio Nobel."},
		{Id: 3, Nome: "Isaac Newton", Historia: "Matemático e físico, formulou as leis do movimento e da gravitação universal."},
		{Id: 4, Nome: "Ada Lovelace", Historia: "Considerada a primeira programadora de computadores, trabalhou com Charles Babbage."},
		{Id: 5, Nome: "Nikola Tesla", Historia: "Inventor e engenheiro, conhecido por suas contribuições ao desenvolvimento da eletricidade e do magnetismo."},
	}

	database.ConectaComBancoDeDados()

	fmt.Println("Starting server on :8080")
	routes.HandleRequests()
}
