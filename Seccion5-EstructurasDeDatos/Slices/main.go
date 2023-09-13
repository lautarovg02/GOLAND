package main

import "fmt"

func main() {
	//Slice declarada
	//var slice []int
	//Slice declarada e inicializada
	diasSemana := []string{"Lunes", "Martes", "Miercoles", "Jueves",
		"Viernes", "Sabado", "Domingo"}
	fmt.Println(diasSemana)
	// A partir de una slice podemos crear otra
	diaSlice := diasSemana[0:3]
	/*
	*Function appende para add elementos a una slice
	*Esta recibe la slice y los elementos para agregarle
	*Y retorna una nueva slice con los elem nuevos
	 */
	diaSlice = append(diaSlice, "Jueves", "Viernes", "Sabado", "Domingo", "Otro dia")
	//Borrar un elemento de una slice
	diaSlice = append(diaSlice[:2], diaSlice[3:]...)
	fmt.Println(diaSlice)
	fmt.Println(len(diaSlice))
	fmt.Println(cap(diaSlice)) // la capacidad es 7 pq se creo a partir de la slice diasSemana

}
