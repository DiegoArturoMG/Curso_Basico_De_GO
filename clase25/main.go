package main

import "fmt"

type figuras2D interface {
	area() float64
}

type cuadrado struct {
	base float64
}

type rectangulo struct {
	base float64
	altura float64
}

func (c cuadrado) area() float64 { 
	return c.base * c.base
}

func (c rectangulo) area() float64 { 
	return c.base * c.altura
}

func calcular(f figuras2D) {
	fmt.Println("Area: ", f.area())
}

func main()  {
	myCuadrado := cuadrado{base: 10}
	myRectangulo := rectangulo{base: 10, altura: 5}
	
	calcular(myCuadrado)
	calcular(myRectangulo)

	// Lista interfaces
	myInterface := []interface{}{"Hola", 12, 4.90}
	fmt.Println(myInterface...)
}