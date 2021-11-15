package main

import "fmt"

func main() {
	func(texto string) {
		fmt.Println(texto)
	}("Olá") // os dois ultimos parenteses servem para executar minha função anonima, e é aqui que eu passo meu parametro já que ela noa podera ser chamada no decorrer do código
}
