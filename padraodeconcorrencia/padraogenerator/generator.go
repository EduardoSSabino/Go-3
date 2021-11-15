package main

import (
	"fmt"
	"time"
)

func main() {
	canal := escrever("Olá mundo!") // um canal do tipo que recebe

	for i := 0; i < 10; i++ {
		fmt.Println(<-canal)
	}
}

func escrever(texto string) <-chan string {
	/*Essa é minha função que vai atender meu padrão generator.
	qual a vantagem dela? Quando eu for chamar a função main,
	eu não vou precisar chamar com a cláusula go, posso chamar
	ela normalmente, como se fosse uma função qualquer. Então
	dentro da função main, vou chamar a função escrever, e como
	ela retorna um canal, vou ter uma variável pra armazenala.*/
	canal := make(chan string)

	go func() {
		for {
			canal <- fmt.Sprintf("Valor recebido: %s", texto)
			time.Sleep(time.Millisecond * 500)
		}
	}()

	return canal
}
