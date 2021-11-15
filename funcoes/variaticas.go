package main

import "fmt"

//nas funcoes variaticas, não precisamos especficar quantos parametros ela irá receber
func soma(numeros ...int) int {
	total := 0
	for _, numero := range numeros {
		total += numero
	}
	return total
}

func escrever(texto string, numeros ...int) { //eesse método recebe mais um tipo de variavel
	for _, numero := range numeros {
		fmt.Println(texto, numero)
	}
}

func main() {
	totalSoma := soma(1, 2, 3, 4, 5, 6, 7, 8)
	fmt.Println(totalSoma)
}
