package main

import (
	"fmt"
	"time"
)

/*Os canais são outra forma de sicronizar duas goroutines (go), é uma maneira mais
sofisticada que o waitgroup (grupo de espera)*/

func main() {
	canal := make(chan string) //um canal que só poderá trafegar dados do tipo string

	go olaMundo("Hello World!", canal)
	fmt.Println("Depos da função olaMundo começar a ser executada")

	/*A diferença entre o canal e oa goroutine, é que o canal tem que esperar receber algo pra prosseguir
	  já a goroutine não precisa, ela não espera a função encerrar pra processeguir, já vai dando continuidade na execução */

	/*CUIDADOS COM OS DEADLOCK!!!
	Mas o qeu seria um deadlock? Deadlock é quando nosso canal ainda está aguardando receber um valor sendo que esse valor
	nunca vai chegar, neste exemplo, minha função olaMundo ira enviar 5x algum valor pro meu canal, porém, meu canal
	está inserido num laço infinitoou seja, SEMPRE estará esperando um valor, isso dará um erro (all goroutines are sleeping), ou seja,
	todas as gorouines estão dormindo, pq estão esperando um valor que nunca irá chegar, fazendo que meu código fique parado*/

	for { //crio um FOR infinito
		mensagem, aberto := <-canal /*Só vou passar dessa linha quando algum valor for enviado pro canal".
		Estou dizendo que meu canal está aguardando receber um valor*/
		if !aberto { //estou tratando meu deadlock
			break //saindo do looping infinito
		}
		fmt.Println(mensagem)
	}
	fmt.Println("Fim do programa!")
}

func olaMundo(texto string, canal chan string) {
	time.Sleep(time.Second * 5)
	for i := 1; i < 6; i++ { //vai enviar um valor 5 vezes pro meu canal
		canal <- texto //meu canal está recebendo esse texto
		time.Sleep(time.Second)
	}

	close(canal) //depois de executar meu FOR 5x, eu vou fechar meu canal. Ele não poderá receber/enviar dados.
}
