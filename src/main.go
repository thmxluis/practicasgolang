package main

import "fmt"

func normalFunction(message string) {
	fmt.Println(message)
}

func tripeArgumentFunction(message string, a, b int) {
	fmt.Println(message, a, b)
}

func returnValue(a int, b int) int {
	return (a + b) * 2
}

func doubleReturnValue(a int) (int, int) {
	return a, a * 2
}

func main() {
	//  Funciones
	fmt.Println("Funciones")
	normalFunction("Hola Mundo")
	tripeArgumentFunction("Hola Mundo", 1, 2)
	fmt.Println(returnValue(1, 2))
	fmt.Println(doubleReturnValue(2))

	// Funciones anónimas que retorna 2 valores
	value1, value2 := doubleReturnValue(2)
	fmt.Println(value1, value2)

	//  Funciones anónimas que extrae 1 solo valor donde retorna 2 valores
	value3, _ := doubleReturnValue(3)
	fmt.Println(value3)

}
