package main

import (
	"fmt"
	"linhadecomando/app"
	"log"
	"os"
)

/*Iremos fazer uma aplicação de duas ações (métodos). O programa receberá um endereço (URL) da web, e e procurar o IP desse endereço,
por exemplo, se passarmos o endereço google.com, ela vai me devolver o IP público e terá uma segunda ação que vai me devolver
o nome do servidor onde o endereço está hospedado. Iremos usar um pacote externo!*/

func main() {
	fmt.Println("Ponto de partida")

	aplicacao := app.Gerar()

	erro := aplicacao.Run(os.Args)
	if erro != nil {
		log.Fatal(erro)
	}
}
