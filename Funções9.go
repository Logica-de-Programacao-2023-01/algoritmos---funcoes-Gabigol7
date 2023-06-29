package main

import (
	"errors"
	"strings"
)

func extractWords(input string) ([]string, error) {
	if len(input) == 0 {
		return nil, errors.New("String vazia")
	}

	words := strings.Split(input, " ")
	return words, nil
}

func main() {
	input := "Olá, mundo! Este é um exemplo de string"
	words, err := extractWords(input)
	if err != nil {
		panic(err)
	}

	for _, word := range words {
		println(word)
	}
}
