package main

import (
	"errors"
	"fmt"
	"math"
)

func isPrime(number int) (bool, error) {
	if number < 0 {
		return false, errors.New("O número deve ser positivo")
	}

	if number < 2 {
		return false, nil
	}

	sqrt := int(math.Sqrt(float64(number)))

	for i := 2; i <= sqrt; i++ {
		if number%i == 0 {
			return false, nil
		}
	}

	return true, nil
}

func main() {
	result, err := isPrime(17)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}
}
