package main

import "errors"

func main() {

}
func numerosComuns(slice1, slice2 []int) ([]int, error) {
	if len(slice1) == 0 || len(slice2) == 0 {
		return nil, errors.New("Um dos slices está vazio")
	}

	numerosComuns := []int{}

	for _, num1 := range slice1 {
		for _, num2 := range slice2 {
			if num1 == num2 {
				numerosComuns = append(numerosComuns, num1)
				break
			}
		}
	}

	return numerosComuns, nil
}
