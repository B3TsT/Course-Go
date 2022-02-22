package main

import "fmt"

func main() {
	var a, b, sum int
	fmt.Print("Ingrese numero: ")
	fmt.Scanln(&a)
	fmt.Print("Ingrese numero")
	fmt.Scanln(&b)
	sum = sumar(a, b)
	fmt.Println("La suma es:", sum)
}

func sumar(a, b int) int {
	return a + b
}
