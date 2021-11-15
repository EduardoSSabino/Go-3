package main

import (
	"fmt"
	"time"
)

//CONCORRENCIA != PARALELISMO

//PARALELISMO: Para isso, teriamos que ter dois processadores no mínimo, paralelismo significa executar duas ações ao mesmo tempo

/*CONCORRÊNCIA: concorrência já se explica somente pelo nome. É quando duas funções são acionadas juntas porém vai ser alternando,
concorrendo/revezando pra ver quem vai ser executada, mas sempre uma de cada vez*/

func main() {
	go olaMundo("Olá mundo!") //goroutine
	olaMundo("Hello world!")
}

func olaMundo(texto1 string) {
	for {
		fmt.Println(texto1)
		time.Sleep(time.Second)
	}
}
