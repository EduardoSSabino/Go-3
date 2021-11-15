package main

import "fmt"

// af unçã init, nada mais é do que uma função que será exxecutada antes da função main

func main() {
	fmt.Println("Função main sendo executada")
}

func init() {
	fmt.Println("Executando a função init")
}
