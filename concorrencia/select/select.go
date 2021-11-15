package main

import (
	"fmt"
	"time"
)

/*SELECT: tem uso bem parecido com switch em questao de sintaxe! porém é de uso exclusivo para concorrência!*/

/*no meu terceiro FOR (na parte em que as linhas estão comentadas), eu vou alterando minha impressão no console entre mensagemCanal1 e mensagemCanal2
mas reparem, meu canal2 tem um time.Sleep de 4 segundo e meu cacal1 tem um time.Sleep de 1 segundo
ou seja, enquanto mensagemCanal2 é impresso duas vezes, mensagemCanal1 deveria ser impresso 4 vezes
mas não é isso que vai acontecer, ele fica travado esperando o canal2 ser executado, pra resolver esse delay desnecessário, iremos usar o select! */

func main() {
	canal1, canal2 := make(chan string), make(chan string)

	go func() { //goroutine que recebe função anonima
		for {
			time.Sleep(time.Second) //espero um segundo
			canal1 <- "Canal 1"     //passo a string pra dentro do meu canal1
		}
	}() //esses parentes executam minha função anonima

	go func() {
		for {
			time.Sleep(time.Second * 4)
			canal2 <- "Canal 2"
		}
	}()

	for {
		/*	mensagemCanal1 := <-canal1  //aguarda canal1 receber um valor pra poder prosseguir
			fmt.Println(mensagemCanal1) //imprimo o que está dentro do meu canal1

			mensagemCanal2 := <-canal2
			fmt.Println(mensagemCanal2)*/

		select {
		case mensagemCanal1 := <-canal1:
			fmt.Println(mensagemCanal1)
		case mensagemCanal2 := <-canal2:
			fmt.Println(mensagemCanal2)
			//PRONTO! acabei com o meu delay desnecaario!
		}

	}
}
