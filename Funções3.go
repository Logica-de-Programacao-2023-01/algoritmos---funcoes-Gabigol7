package main

func main() {

}
func concatenarStrings(slice []string) string {
	resultado := ""
	for _, str := range slice {
		resultado += str
	}
	return resultado
}
