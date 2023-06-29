package main

import (
	"errors"
	"fmt"
	"strings"
)

func replaceVowels(str string) (string, error) {
	if str == "" {
		return "", errors.New("String vazia")
	}

	vowels := []string{"a", "e", "i", "o", "u", "A", "E", "I", "O", "U"}

	for _, vowel := range vowels {
		str = strings.ReplaceAll(str, vowel, "*")
	}

	return str, nil
}

func main() {
	// Exemplo de uso
	newString, err := replaceVowels("Hello World")
	if err != nil {
		panic(err)
	}
	fmt.Println(newString) // Output: H*ll* W*rld
}
