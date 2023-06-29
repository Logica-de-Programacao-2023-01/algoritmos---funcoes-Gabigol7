package main

import "fmt"

func main() {

}
func applyFunctionToSlice(slice []int, fn func(int) int) ([]int, error) {
	if len(slice) == 0 {
		return nil, fmt.Errorf("O slice está vazio")
	}

	result := make([]int, len(slice))
	for i, value := range slice {
		result[i] = fn(value)
	}

	return result, nil
}
