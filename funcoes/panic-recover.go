package main

import "fmt"

func recuperarExecucao() { //RECOVER
	r := recover()
	if r != nil { //FAZENDO O TRATAMENTO DO PANIC
		fmt.Println("Execução recuperada com sucesso!")
	}
}

func aprovacaoAluno(notas int) bool {
	defer recuperarExecucao()
	if notas < 6 {
		return false
	} else if notas > 6 {
		return true
	}

	//SITUAÇÃO HIPOTÉTICA! O único resultado que não é cabível as situações acima, é se o valor for igual a 6, supondo que em hipotese alguma a nota possa ser 6, usaremos o panic

	panic("A MÉDIA É EXATAMENTE 6!")
	//Quando a função panic é acionada, ela para tudo e chamar os defer pra tentar recuperar o código, e se caso não resolver, acoina o recover mantém tudo parado!
}

func main() {
	fmt.Println(aprovacaoAluno(6))
	fmt.Println("Pós execução.")
}
