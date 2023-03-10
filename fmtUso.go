package main

import (
	"fmt"
)

func main() {
	// Declaracion de variables
	helloMessage := "Hello"
	worldMessage := "World"

	// Println
	fmt.Println(helloMessage, worldMessage)
	fmt.Println(helloMessage, worldMessage)
	// Printf
	nombre := "Platzi"
	cursos := 500
	fmt.Printf("%s tiene mas de %d cursos\n", nombre, cursos)
	fmt.Printf("%v tiene mas de %v cursos\n", nombre, cursos)
	// Sprintf
	message := fmt.Sprintf("%v tiene mas de %v cursos", nombre, cursos)
	fmt.Println(message)
	// Tipo datos
	fmt.Printf("helloMessage: %T\n", helloMessage)
	fmt.Printf("cursos: %T\n", cursos)

}