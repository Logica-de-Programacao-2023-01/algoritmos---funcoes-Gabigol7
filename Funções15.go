package main

import "errors"

func main() {

}
func verificarPresenca(numero int, slice []int) (bool, error) {
	if len(slice) == 0 {
		return false, errors.New("Slice vazio")
	}

	for _, valor := range slice {
		if valor == numero {
			return true, nil
		}
	}

	return false, nil
}
