package main

import (
	"fmt"
	"math"
)

type forma interface {
	area() float64
}

func escreverArea(f forma) {
	fmt.Printf("A area da forma é de %0.2f\n", f.area())
}

func (retangulo1 retangulo) area() float64 { //a assinatura tem que ser identica a da interface!
	return retangulo1.altura * retangulo1.largura
}

func (circulo1 *circulo) area() float64 {
	return math.Pi * math.Pow(circulo1.raio, 2)
}

type retangulo struct {
	altura  float64
	largura float64
}

type circulo struct {
	raio float64
}

func main() {
	retangulo1 := retangulo{altura: 10.0, largura: 5.0}
	escreverArea(retangulo1)
	circulo1 := circulo{raio: 100.0}
	escreverArea(&circulo1)
}
