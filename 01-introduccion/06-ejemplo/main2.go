package main

import "fmt"

func main() {
	var a, b, core, menu int
	fmt.Print("ingrese numero: ")
	fmt.Scanln(&a)
	fmt.Print("ingrese numero: ")
	fmt.Scanln(&b)
	fmt.Print("Opcion de menu:\n1.Calcular conciente\n2.Calcular Residuo\n")
	fmt.Scanln(&menu)
	switch menu {
	case 1:
		core = cociente(a, b)
		fmt.Println("Resultado del cociente: ", core)
	case 2:
		core = residuo(a, b)
		fmt.Println("Resultado del residuo", core)
		break
	}
}
func cociente(a, b int) int {
	return a / b
}
func residuo(a, b int) int {
	return a % b
}
