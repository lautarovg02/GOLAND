package main

import (
	"fmt"
	"time"
)

func main() {

	// lo que hace la funcion Now del Paquete time es decirnos la hora actual
	t := time.Now()
	hora := t.Hour()

	if hora < 12 {
		fmt.Println("Manana")
	} else if hora < 17 {
		fmt.Println("Tarde")
	} else {
		fmt.Println("Noche")
	}
}
