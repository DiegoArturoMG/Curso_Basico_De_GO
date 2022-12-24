package main

import "fmt"

func main() {
	switch modulo := 5 % 2; modulo {
	case 0:
		fmt.Println("Es par")
	default:
		fmt.Println("Es impar")
	}

	// Sin condicion
	value := 50
	switch {
	case value > 100:
		fmt.Println("Mayor a 100")
	case value < 0:
		fmt.Println("Menor a 0")
	default:
		fmt.Println("No condicion")
	}
}