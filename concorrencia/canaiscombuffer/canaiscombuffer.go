package main

import "fmt"

func main() {
	canal := make(chan string, 2) //crio um canal com buffer de capacidade 2
	canal <- "Olá mundo!"
	canal <- "Hello World" //até aqui tá tranquilo, o problema é causa mais um deadlock, ai sim estaria excendendo a capacidade do meu buffer

	mensagem := <-canal //estou aguardando enviarem um valor pro canal pra ai sim continuar e ir pra próxima linha
	mensagem2 := <-canal

	fmt.Println(mensagem)
	fmt.Println(mensagem2)
}
