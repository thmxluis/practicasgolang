package main

import (
	"log"
	"strconv"
)

func main() {

	// condicional if

	valor := 20

	if valor == 10 {
		println("El valor es 10")
	}

	if valor := 10; valor == 10 {
		println("El valor es 10")
	}

	// condicional switch

	switch valor {
	case 10:
		println("El valor es 10")
	case 20:
		println("El valor es 20")
	default:
		println("El valor es otro")
	}

	// condicional  with  and	& or

	if valor > 10 && valor < 30 {
		println("El valor es mayor a 10 y menor a 30")
	}

	if valor > 10 || valor < 30 {
		println("El valor es mayor a 10 o menor a 30")
	}

	// convertir texto a numero
	value, err := strconv.Atoi("213123")
	if err != nil {
		println("Error al convertir")
		log.Fatal(err)
	} else {
		println(value)
	}

}
