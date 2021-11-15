package main

import "fmt"

func closure() func() {
	texto := "Dentro da função closure"

	funcao := func() {
		fmt.Println(texto)
	}
	return funcao
}

//vou chamar a função closure dentro do main

func main() {
	texto := "Dentro da função main"
	fmt.Println(texto)

	funcaoNova := closure()
	funcaoNova()
	/*a função closure tem uma espécie de memória, neste caso a pergunta que fica é: qual "texto" minha funcaoNova irá imprimir?
	o "texto" da func main ou da func closure? da closure!*/
}
