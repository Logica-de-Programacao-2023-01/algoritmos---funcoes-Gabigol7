package main

import (
	"errors"
	"strings"
)

func main() {

}
func concatenateStrings(slice []string) (string, error) {
	if len(slice) == 0 {
		return "", errors.New("O slice está vazio")
	}

	result := strings.Join(slice, ", ")
	return result, nil
}


	stringsSlice := []string{"Hello", "world", "OpenAI"}
	concatenated, err := concatenateStrings(stringsSlice)
	if err != nil {
		// Tratar o erro
	} else {
		// Utilizar a string concatenada
	}