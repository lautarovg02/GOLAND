package main

import (
	"fmt"
	"strconv"
)

func main() {
	var integer16 int16 = 50
	var integer32 int32 = 100
	var (
		f32  float32 = 43.21
		f64  float64 = 244324232.21
		ui8  uint8   = 3
		ui16 uint16  = 133
	)
	/*
	*Como son variables de distinto tipo, para sumarlas
	* tenemos que realizar una conversion, podemos usar funciones integradas para eso como: int32() o int16(),
	* colocando adentro el tipo de dato que queremos convertir, lo mismo para float o uint
	 */
	fmt.Println("La suma essss\t", int32(integer16)+integer32)
	fmt.Println("La suma essss\t", uint16(ui8)+ui16)
	fmt.Println("La suma essss\t", float64(f32)+f64)

	/*
	*Conversiones mas Complejas, String a int
	*Vamos a utilizar un paquete strconv
	*La funcion encargada de convertir de string a int es Atoi
	 */
	s := "20"
	/*
	*Cuando esperamos valores de una funcion
	* sino queremos tomar uno lo remplazamaos con un _
	 */
	deStringAEntero, _ := strconv.Atoi(s)
	fmt.Println(deStringAEntero)

	/*
	*Bien si ahora queremos convertir de forma viceversa(entero a cadena)
	*
	 */

	n := 42
	deEnteroAString := strconv.Itoa(n)
	fmt.Println(deEnteroAString)

}
