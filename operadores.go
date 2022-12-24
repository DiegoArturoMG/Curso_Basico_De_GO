package main

import (
	"fmt"
	"math"
)

func main() {
	// Area del Rectangulo
	base := 4
	altura := 2
	fmt.Println("Area Rectangulo:", base * altura )

	// Area Trapecio
	base1 := 2
	base2 := 4
	areaTrapecio := altura * (base1 + base2) / 2
	fmt.Println("Area Trapecio:", areaTrapecio)

	// Area Circulo
	radio := 2
	areaCirculo := math.Pi * math.Pow(float64(radio), 2)
	fmt.Println("Area Circulo:", areaCirculo)
}