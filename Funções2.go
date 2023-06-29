package main

import (
	"fmt"
	"strings"
)

func contarVogais(texto string) int {
	vogais := "aeiouAEIOU"
	count := 0

	for _, char := range texto {
		if strings.ContainsRune(vogais, char) {
			count++
		}
	}

	return count
}

func main() {
	texto := "Hello, World!"
	quantidade := contarVogais(texto)
	fmt.Printf("A quantidade de vogais em '%s' é: %d\n", texto, quantidade)
}
