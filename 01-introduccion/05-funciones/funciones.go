package main

import "fmt"

func main() {
	saludar("juan")
	r := suma(4, 7)
	fmt.Println(r)
}

func saludar(nombre string) {
	fmt.Println("Hola", nombre)
}

func suma(a, b int) int {
	return a + b
}
