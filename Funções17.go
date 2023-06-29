package main

import (
	"errors"
	"sort"
	"strings"
)

func ordemAlfabetica(slice []string) (string, error) {
	if len(slice) == 0 {
		return "", errors.New("O slice está vazio")
	}

	sort.Strings(slice)
	result := strings.Join(slice, " ")

	return result, nil
}
