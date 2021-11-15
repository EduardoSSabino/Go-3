package main

import "fmt"

func funcao1() {
	fmt.Println("Executando função 1!")
}

func funcao2() {
	fmt.Println("Executando função 2!")
}

func aprovacaoDoAluno(notas int) bool {
	if notas < 6 {
		return false
	}
	return true
}
func main() {
	defer funcao1() //o uso do defer faz com que a linha seja executada por ultimo, independete de qualque coisa!
	// DEFER = ADIAR
	funcao2()
	fmt.Println(aprovacaoDoAluno(8))
}
