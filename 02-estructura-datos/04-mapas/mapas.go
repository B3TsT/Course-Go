package main

import "fmt"

func main() {
	dias := make(map[int]string)
	fmt.Println(dias)
	//Agregar datos
	dias[1] = "domingo"
	dias[2] = "lunes"
	dias[3] = "martes"
	dias[4] = "miercoles"
	dias[5] = "jueves"
	dias[6] = "viernes"
	dias[7] = "sabado"
	fmt.Println(dias)
	dias[7] = "SABADO"
	fmt.Println(dias)
	delete(dias, 1)
	fmt.Println(dias, len(dias))

	//Nueva mapa
	estudiantes := make(map[string][]int)
	estudiantes["Fran"] = []int{12, 15, 16}
	estudiantes["roel"] = []int{13, 45, 12}
	fmt.Println(estudiantes)
	fmt.Println(estudiantes["Fran"][1])
}
