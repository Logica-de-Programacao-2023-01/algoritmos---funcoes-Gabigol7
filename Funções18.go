package main

import "fmt"

func main() {

}
func applyFunction(slice []int, fn func(int) int) (int, error) {
	if len(slice) == 0 {
		return 0, fmt.Errorf("O slice está vazio")
	}

	sum := 0
	for _, num := range slice {
		result := fn(num)
		sum += result
	}

	return sum, nil
}
