package main

import "fmt"

/*Quando faemos uma alteração sem usar o ponteiro, a gente ta alterando um cópia daquele balor,
quando alteramos usando o ponteiro, estamos mudando a "variavel original", mudamos ela no código todo"*/

func inverterSinal(numero int) int {
	return numero * -1
}

//Resumidamente, funções com ponterios, a gente mexe no endereço de memória e não no valor

func inverterSinalComPonteiro(numero *int) {
	*numero = *numero * -1 //estou jogando um valor novo dentro do endereço de memória
}
func main() {
	numero := 20
	numeroInvertido := inverterSinal(numero)
	fmt.Println(numeroInvertido)

	novoNumero := 40
	fmt.Println(novoNumero)
	inverterSinalComPonteiro(&novoNumero)
	fmt.Println(novoNumero) //imprimi meu número pós alteração no endereço de memória
}
