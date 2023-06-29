package main

import "errors"

func main() {

}
func calcularMedia(valores []int) (float64, error) {
	if len(valores) == 0 {
		return 0, errors.New("O slice está vazio")
	}

	soma := 0
	for _, valor := range valores {
		soma += valor
	}

	media := float64(soma) / float64(len(valores))
	return media, nil
}
