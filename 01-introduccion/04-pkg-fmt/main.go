package main

import "fmt"

func main() {
	hola := "hola"
	mundo := "mundo"

	fmt.Println(hola)
	fmt.Println(mundo)

	nombre := "alex"
	edad := 25

	fmt.Printf("hola, %s y su nombre es %d \n", nombre, edad)

	fmt.Printf("hola, %v y su edad es %v \n", nombre, edad)

	mensaje := fmt.Sprintf("hola %s y su edad es %d", nombre, edad)

	fmt.Println(mensaje)

	fmt.Printf("nombre: %T \n", edad)

	fmt.Print("Ingrese nombre:")
	fmt.Scanln(&nombre)
	fmt.Println("otro nombre: ", nombre)
}
