package main

import "errors"

func main() {

}
func filterStrings(slice []string) ([]string, error) {
	if len(slice) == 0 {
		return nil, errors.New("Slice está vazio")
	}

	filteredSlice := make([]string, 0)
	for _, str := range slice {
		if len(str) > 5 {
			filteredSlice = append(filteredSlice, str)
		}
	}

	return filteredSlice, nil
}
