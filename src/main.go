package main

import "fmt"

func main() {
	// // Declaracion de constantes
	// const pi float64 = 3.14
	// const pi2 = 3.1416

	// // Declaracion de variables enteras
	// base := 12
	// var altura int = 14
	// var area int

	// // Zero values
	// var a int
	// var b float64
	// var c string
	// var d bool

	// fmt.Println("pi: ", pi2)
	// fmt.Println("pi2: ", pi2)
	// fmt.Println("base: ", base)
	// fmt.Println("altura: ", altura)
	// fmt.Println("area: ", area)
	// fmt.Println("a: ", a)
	// fmt.Println("b: ", b)
	// fmt.Println("c: ", c)
	// fmt.Println("d: ", d)

	// Area cuadrado
	const baseCuadrado = 10
	areaCuadrado := baseCuadrado * baseCuadrado
	fmt.Println("areaCuadrado: ", areaCuadrado)

	// Suma
	x := 10
	y := 50
	result := x + y
	fmt.Println("result: ", result)

	// Resta
	x2 := 10
	y2 := 50
	result2 := y2 - x2
	fmt.Println("result2: ", result2)

	// Multiplicacion
	x3 := 10
	y3 := 50
	result3 := x3 * y3
	fmt.Println("result3: ", result3)

	// Division
	x4 := 60
	y4 := 50
	result4 := x4 / y4
	fmt.Println("result4: ", result4)

	// Modulo
	x5 := 60
	y5 := 50
	result5 := x5 % y5
	fmt.Println("result5: ", result5)

	// Incremento
	x6 := 10
	x6++
	fmt.Println("x6: ", x6)

	// Decremento
	x7 := 10
	x7--
	fmt.Println("x7: ", x7)

	// Area rectangulo
	const baseRectangulo = 10
	const alturaRectangulo = 50
	areaRectangulo := baseRectangulo * alturaRectangulo
	fmt.Println("areaRectangulo: ", areaRectangulo)

	// Area triangulo
	const baseTriangulo = 10
	const alturaTriangulo = 50
	areaTriangulo := (baseTriangulo * alturaTriangulo) / 2
	fmt.Println("areaTriangulo: ", areaTriangulo)

	// Area circulo
	const pi = 3.14
	const radio = 10
	areaCirculo := pi * (radio * radio)
	fmt.Println("areaCirculo: ", areaCirculo)

}
