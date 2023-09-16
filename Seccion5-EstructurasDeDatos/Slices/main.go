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
	/*
	*Borrar un elemento de una slice
	*append(diaSlice[:2], diaSlice[3:]...)
	 */
	diaSlice = append(diaSlice[:2], diaSlice[3:]...)
	fmt.Println(diaSlice)
	fmt.Println(len(diaSlice))
	fmt.Println(cap(diaSlice)) // la capacidad es 7 pq se creo a partir de la slice diasSemana

	/*Function make([]tipo de dato, longitud, capacidad)
	*Con eesta funcion podemos crear una slice
	 */
	nombres := make([]string, 5, 10)
	nombres[0] = "Lau"
	nombres = append(nombres, "Felipe", "Josefina")
	fmt.Println(nombres)

	//Function copy podemos copiar elementos de una slice a otra
	slice1 := []int{1, 2, 3, 4, 5}
	slice2 := make([]int, 5)
	fmt.Println(copy(slice2, slice1))
	fmt.Println(slice1)
	fmt.Println(slice2)
}
