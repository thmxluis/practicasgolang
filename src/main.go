package main

import "fmt"

func main() {

	m := make(map[string]int)

	m["jose"] = 14
	m["pepito"] = 20
	fmt.Println(m)

	// Recorrer el map
	for k, v := range m {
		fmt.Println(k, v)
	}

	// Encontrar un valor
	v, ok := m["jose"]
	fmt.Println(v, ok)

}
