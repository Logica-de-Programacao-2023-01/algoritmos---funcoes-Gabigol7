package main

import (
	"errors"
	"sort"
)

func main() {

}
func ordenarSlice(slice []int) ([]int, error) {
	if len(slice) == 0 {
		return nil, errors.New("O slice está vazio")
	}

	// Cria uma cópia do slice original para não modificar o slice original
	// durante a ordenação
	copiaSlice := make([]int, len(slice))
	copy(copiaSlice, slice)

	// Ordena o slice em ordem crescente
	sort.Ints(copiaSlice)

	return copiaSlice, nil
}
