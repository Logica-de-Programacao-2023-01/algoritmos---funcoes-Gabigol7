package main

import (
	"errors"
	"math"
)

// Verifica se um número é primo
func isPrime(num int) bool {
	if num < 2 {
		return false
	}
	for i := 2; i <= int(math.Sqrt(float64(num))); i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}

// Retorna um slice com todos os números primos menores ou iguais ao número fornecido
func getPrimes(num int) ([]int, error) {
	if num < 0 {
		return nil, errors.New("O número deve ser positivo")
	}

	primes := []int{}
	for i := 2; i <= num; i++ {
		if isPrime(i) {
			primes = append(primes, i)
		}
	}

	return primes, nil
}

func main() {
	primes, err := getPrimes(10)
	if err != nil {
		panic(err)
	}

	// Imprime os números primos
	for _, prime := range primes {
		println(prime)
	}
}
