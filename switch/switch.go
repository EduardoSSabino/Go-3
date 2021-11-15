package main

import "fmt"

func maioridade(idade int) string {
	switch {
	case idade < 18:
		return "Você é menor de idade!"
	default:
		return "Você é maior de idade!"
	}
}

//outra maneira de fazer o switch

func diaDaSemana(dia int) string {
	switch dia {
	case 1:
		return "Domingo"
	case 2:
		return "Segunda-feira"
	case 3:
		return "Terça-feira"
	case 4:
		return "Quarta-feira"
	case 5:
		return "Quinta-feira"
	case 6:
		return "Sexta-feira"
	case 7:
		return "Sábado"
	default:
		return "Número inválido"
	}
}
func main() {
	fmt.Println("Switch")

	var diaDeSemana, maioridadeUsuario int

	fmt.Println("Digite um dia da semana: ")
	fmt.Scanf("%d\n", &diaDeSemana)

	fmt.Println("Digite a sua idade: ")
	fmt.Scanf("%d\n", &maioridadeUsuario)

	fmt.Println(diaDaSemana(diaDeSemana))

	fmt.Println(maioridade(maioridadeUsuario))
}
