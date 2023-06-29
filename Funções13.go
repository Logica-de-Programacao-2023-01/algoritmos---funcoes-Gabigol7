package main

import (
	"errors"
	"fmt"
)

func somarDigitos(numero int) (int, error) {
	if numero < 0 {
		return 0, errors.New("O número não pode ser negativo")
	}

	soma := 0
	for numero != 0 {
		digito := numero % 10
		soma += digito
		numero /= 10
	}

	return soma, nil
}

func main() {
	numero := 12345
	soma, err := somarDigitos(numero)
	if err != nil {
		fmt.Println("Erro:", err)
	} else {
		fmt.Println("A soma dos dígitos de", numero, "é:", soma)
	}
}
