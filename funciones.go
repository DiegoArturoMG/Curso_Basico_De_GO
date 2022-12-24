package main

import "fmt"

func normalFunction(message string) {
	fmt.Println(message)
}

func tripleArg(a int, b int, c string) {
	fmt.Println(a, b, c)
}

func returnValue(a int) int {
	return a * 2
}

func doubleReturn(a int) (c, d int) {
	return a, a * 2
}

func main() {
	normalFunction("Hola Mundo")
	tripleArg(1, 2, "hola")

	value := returnValue(2)
	fmt.Println("Value:", value)

	value1, value2 := doubleReturn(2)
	fmt.Println(value1, value2)

	_, value3 := doubleReturn(2)
	fmt.Println(value3)
}