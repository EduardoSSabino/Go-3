package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type cachorro struct {
	Nome  string `json:"nome"`
	Raca  string `json:"raca"`
	Idade uint   `json:"idade"`
}

func main() {
	cachorroEmJSON := `{"nome":"Rex","raca":"Dálmata","idade":3}` //formato json
	var c cachorro
	erro := json.Unmarshal([]byte(cachorroEmJSON), &c)
	if erro != nil {
		log.Fatal(erro)
	}

	fmt.Println(c)

	cachorro2EmJSON := `{"nome":"Preciosa","raca":"Basset}`

	c2 := make(map[string]string)
	erro2 := json.Unmarshal([]byte(cachorro2EmJSON), &c2)
	if erro2 != nil {
		log.Fatal(erro2)
	}
	fmt.Println(c2)
}
