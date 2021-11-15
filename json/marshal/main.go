package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
)

/*JSON é um método que usanos pra trafegar dados. */

type cachorro struct {
	Nome  string `json:"nome"` //mostro qual item do json meu atributo do struct irá representar
	Raca  string `json:"raca"`
	Idade uint   `json:"idade"`
}

func main() {

	cachorro1 := cachorro{Nome: "George", Raca: "Chow-Chow", Idade: 5}
	fmt.Println(cachorro1)

	cachorroEmJSON, erro := json.Marshal(cachorro1) //Converte um método ou uma struct para json
	if erro != nil {
		log.Fatal(erro) //se tiver um erro, pare o programa
	}
	fmt.Println(cachorroEmJSON) //aqui eu imprimo um slice de números, usarei o coomadno a baixo pra converter pra JSON

	fmt.Println(bytes.NewBuffer(cachorroEmJSON)) //extensão necessária pra podermos ver o arquivo em JSON

	//json.Unmarshal()//Passa um json para um método

	cachorro2 := map[string]string{
		"nome": "Toby",
		"raca": "Poodle",
	}

	cachorro2EmJSON, erro := json.Marshal(cachorro2)
	if erro != nil {
		log.Fatal(erro)
	}

	fmt.Println(cachorro2EmJSON)
	fmt.Println(bytes.NewBuffer(cachorro2EmJSON))
}
