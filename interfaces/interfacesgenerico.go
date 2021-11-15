package main

import "fmt"

func generica(interf interface{}) { //qualquer coisa atende a essa interface, por isso chamamos de interface generica
	fmt.Println(interf)
}

func main() {
	generica("String")
	generica(1)
	generica(true)

	fmt.Println(1, 2, "String", false, true, float64(1234))

	//Não é muito recomendado usar a interface generica, abrimos mão de boa parte da segunrança ao implementar esse tipo de funcionalidade.
}
