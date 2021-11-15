package main

import (
	"fmt"
	"sync"
	"time"
)

//CONCORRENCIA != PARALELISMO

//PARALELISMO: Para isso, teriamos que ter dois processadores no mínimo, paralelismo significa executar duas ações ao mesmo tempo

/*CONCORRÊNCIA: concorrência já se explica somente pelo nome. É quando duas funções são acionadas juntas porém vai ser alternando,
concorrendo/revezando pra ver quem vai ser executada, mas sempre uma de cada vez*/

func main() {
	var waitGroup sync.WaitGroup //irá realizar uma imortação chamada "sync"
	waitGroup.Add(2)             //dentro dos parenteses eu passo a qauntidade de goroutines que estarão sendo acionada ao mesmo tempo

	go func() {
		olaMundo("Olá mundo!")
		waitGroup.Done()
	}()

	go func() {
		olaMundo("Hello world!")
		waitGroup.Done() //ele retirar um do waitGroup.Add pa poder ser executado
	}()

	waitGroup.Wait() //espera minha contagem chegar a zero, ou seja, minhas duas waitGroup ser acionadas!

}

func olaMundo(texto string) {
	for i := 0; i < 6; i++ {
		fmt.Println(texto)
		time.Sleep(time.Second)
	}
}
