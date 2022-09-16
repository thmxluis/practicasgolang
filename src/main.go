package main

import "fmt"

func main() {

	// declare a variable
	helloMessage := "Hello"
	worldMessage := "World!"

	// print the variable

	// println	- print with a new line
	fmt.Println(helloMessage, worldMessage)
	fmt.Println(helloMessage, worldMessage)

	// printf	- print with a format
	nombre := "Juan"
	edad := 20
	fmt.Printf("Hola, mi nombre es %s y tengo %d añosa\n", nombre, edad)
	fmt.Printf("Hola, mi nombre es %v y tengo %v añosa\n", nombre, edad)

	// Sprintf	- print with a format and return a string
	message := fmt.Sprintf("Hola, mi nombre es %s y tengo %d añosa", nombre, edad)
	fmt.Println(message)

	//  tipos de datos
	fmt.Printf("helloMessage: %T\n", helloMessage)
	fmt.Printf("edad: %T\n", edad)

}
