package main

import (
	"fmt"
	"time"
)

func main() {

	// En GO no hay que poner la condicion entre parentesis
	// Y se pueden declarar variables dentro de un if.
	// lo que hace la funcion Now del Paquete time es decirnos la hora actual

	//Podemos hacerlo asi
	if t := time.Now(); t.Hour() < 12 {
		fmt.Println("Manana")
	} else if t.Hour() < 17 {
		fmt.Println("Tarde")
	} else {
		fmt.Println("Noche")
	}

	// o asi

	//t := time.Now();
	//hour := t.Hour()
	//
	//// o asi
	//if hour < 12 {
	//	fmt.Println("Manana")
	//} else if hour < 17 {
	//	fmt.Println("Tarde")
	//} else {
	//	fmt.Println("Noche")
	//}
}
