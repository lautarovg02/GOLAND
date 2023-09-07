package main

//* Para importar varias paquetes usamos el import mas parentesis, como abajo.
import (
	"fmt"
	"rsc.io/quote"
)

//import "fmt"

func main() {
	fmt.Println("Hola mundo")
	//utilizando quote- un paquete de tercero.
	fmt.Println(quote.Hello())
	fmt.Println(quote.Glass())
}
