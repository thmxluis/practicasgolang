package main

import (
	"fmt"
	"strings"
)

func isPalindromo(palindromo string) string {
	palindromo = strings.ToLower(palindromo)
	for i := 0; i < len(palindromo)/2; i++ {
		if palindromo[i] != palindromo[len(palindromo)-1-i] {
			return "No es palindromo"
		}
	}
	return "Es un palindromo"
}

func main() {

	// Slice with range
	slice := []string{"hola", "que", "tal"}

	//
	for i, v := range slice {
		fmt.Println(i, v)
	}
	for _, v := range slice {
		fmt.Println(v)
	}
	for i := range slice {
		fmt.Println(i)
	}

	// Palindromo
	palindromo := "Ama"
	fmt.Println(isPalindromo(palindromo))

}
