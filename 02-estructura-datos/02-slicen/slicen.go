package main

import "fmt"

func main() {
	//slicen
	//diferencias entre array y slicen
	//los slicen son mutables
	//se puede agregar mas longitud/elementos a un slicen que ya esta definido
	numeros := []int{1, 2, 3}
	fmt.Println(numeros, len(numeros))

	//Agrefar datos
	numeros = append(numeros, 4)
	numeros = append(numeros, 5)
	fmt.Println(numeros, len(numeros))

	//subslicen
	subSilcen := numeros[:2]
	numeros[0] = 100
	fmt.Println(subSilcen)
	fmt.Println(numeros)

	//punteos donde inicia y donde termina

	//longitud
	//capacidad

	meses := []string{"enero", "febrero", "marzo"}
	fmt.Printf("Len: %v\nCap: %v\n%p\n", len(meses), cap(meses), meses)
	meses = append(meses, "abril")
	fmt.Printf("Len: %v\nCap: %v\n%p\n", len(meses), cap(meses), meses)
}
