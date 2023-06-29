package main

import (
	"errors"
	"fmt"
)

func obterNumerosPares(slice []int) ([]int, error) {
	if len(slice) == 0 {
		return nil, errors.New("O slice está vazio")
	}

	var numerosPares []int

	for _, num := range slice {
		if num%2 == 0 {
			numerosPares = append(numerosPares, num)
		}
	}

	return numerosPares, nil
}

func main() {
	// Exemplo de uso
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	numerosPares, err := obterNumerosPares(slice)
	if err != nil {
		// Tratar o erro, se necessário
	} else {
		// Utilizar o novo slice com os números pares
		fmt.Println(numerosPares) // Saída: [2 4 6 8 10]
	}
}
