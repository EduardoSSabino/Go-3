package main

import "fmt"

/*É como se eu tivesse uma fila de tarefas a serem executadas, e você tem vários works (processos) pegando itens dessa fila de maneira independente*/

//Supondo que essa minha fila fosse vários números para serem calculados na minha sequencia de fibonacci

func main() {
	tarefas := make(chan int, 45)    //esse canal terá a sequencia de números que precisaremos calcular
	resultados := make(chan int, 45) //e esse canal irá armazenar os resultados na medida que forem calculados

	//quanto mais go worker eu tiver, mais rápido minha leitura será!
	go worker(tarefas, resultados)
	go worker(tarefas, resultados)
	go worker(tarefas, resultados)
	go worker(tarefas, resultados)
	go worker(tarefas, resultados)
	go worker(tarefas, resultados)

	for i := 1; i < 46; i++ {
		tarefas <- i
	}
	close(tarefas) //fecho meu canal de tarefas para que ele não possa receber mais nenhum tipo de dado

	//Agora vou imprimir os resultados na tela

	for i := 1; i < 46; i++ {
		resultado := <-resultados
		fmt.Println(resultado)
	}
}

func worker(tarefas <-chan int, resultados chan<- int) { /*significado das setas: a primeira seta indica que "tarefas" só receberão dados
	e a segunda seta indica que "resultados" só enviará dados*/
	for numero := range tarefas {
		resultados <- fibonacci(numero)
	}
}

func fibonacci(posicao int) int {
	if posicao <= 1 { //minha condição de parada!
		return posicao
	}
	//especifica a condição de parada
	return fibonacci(posicao-2) + fibonacci(posicao-1)
}
