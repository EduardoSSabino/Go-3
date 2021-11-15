package main

import (
	"fmt"
	"math/rand"
	"time"
)

/*A ideia principal do multiplexador é você pegar dois ou mais canais e juntar com um canal só*/

func main() {
	/*Irei chamar a função escrever mais de uma vez, ou seja, vou ter mais de uma canal
	  a ideia é, pegar esse dois canais e juntar em um só*/

	canal := multiplexar(escrever("Olá mundo!"), escrever("Hello world!")) //esout passando dois canais e reornadno um só

	for i := 0; i < 10; i++ {
		fmt.Println(<-canal)
	}
}

func multiplexar(canalDeEntrada1, canalDeEntrada2 <-chan string) <-chan string { //recebe dois canais que só recebem valor, e retorna um canal de string
	canalDeSaida := make(chan string)

	go func() {
		for {
			select {
			case mensagem1 := <-canalDeEntrada1: //pego minha mensagem e jogo pro canal de saída
				canalDeSaida <- mensagem1 //canal saída reebe mensagem
			case mensagem2 := <-canalDeEntrada2:
				canalDeSaida <- mensagem2
			}
		}
	}()
	return canalDeSaida
}

func escrever(texto string) <-chan string {

	canal := make(chan string)

	go func() {
		for {
			canal <- fmt.Sprintf("Valor recebido: %s", texto)
			time.Sleep(time.Millisecond * time.Duration(rand.Intn(2000))) //irá espera um tempo aleatório de 1 a 2000.
		}
	}()

	return canal
}
