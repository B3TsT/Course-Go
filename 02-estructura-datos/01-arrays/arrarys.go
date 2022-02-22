package main

import (
	"fmt"
)

func main() {
	var numeros [5]int

	numeros[0] = 10
	numeros[1] = 20
	numeros[2] = 30
	fmt.Println(numeros)
	fmt.Println(numeros[4])

	//Arrays con valores
	nombres := [3]string{"fran", "luis", "juan"}
	fmt.Println(nombres)

	//Arrays sin longitud
	color := [...]string{
		"rojo",
		"verde",
		"negro",
		"azul",
	}
	fmt.Println(color, len(color))
	//indices definidos
	monedas := [...]string{
		0: "dolar",
		2: "soles",
		3: "pesos",
		5: "euro",
	}
	fmt.Println(monedas, len(monedas))

	//sub array
	//Si se quiere obtener desde el indice 0 no es necesario agregar el primer valor ejemplo: 0:3 -> :3
	subMoneda := monedas[1:3]
	fmt.Println(subMoneda)
	//ultimo valor
	subMoneda2 := monedas[3:]
	fmt.Println(subMoneda2)
}
