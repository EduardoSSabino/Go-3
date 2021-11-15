package main

import "fmt"

type usuario struct { //Atributos
	nome   string
	idade  uint8
	altura float64
}

func (u usuario) salvar(metodo string) { //estou criando um método para minha struct usuario
	fmt.Printf("Dentro do %s salvar!\n", metodo)
}

func (u *usuario) fazerAniversario() { /*estou fazendo uma alteração no endereço de memória,
	ou seja a alteração do valor, será para o programa todo e nao só dentro do escopo da função*/
	u.idade++
}

func main() {
	usuario1 := usuario{nome: "João", idade: 18, altura: 1.80}
	fmt.Println(usuario1)

	usuario1.salvar("método")
	usuario1.fazerAniversario()

	fmt.Printf("O usuario fará %d anos de idade", usuario1.idade)
}
