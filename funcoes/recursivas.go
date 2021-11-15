package main

import "fmt"

//OBS: A sequência de Fibonacci é: o próximo número da sequência é igual a soma dos dois numeros anteriores

//Sequêmcia de Fibonacci: 1, 1, 2, 3, 5, 8, 13...

//São funcçoes que chamam elas mesmas
func fibonacci(posicao uint) uint {
	if posicao <= 1 { //minha condição de parada!
		return posicao
	}
	//especifica a condição de parada
	return fibonacci(posicao-2) + fibonacci(posicao-1)
}
func main() {
	fmt.Println("Funções recursivas!")

	posicao := uint(10) //me trás os 10 primerios valores da sequencia de fibonacci

	for i := uint(0); i < posicao; i++ {
		fmt.Println(fibonacci(i)) //vai escrever todos os números da sequencia de fibonacci
	}

}
