package main

import (
	"fmt"
	"math"
)

func main() {
	/*
	*Datos Numericos
	* todos los tipos de int
	* int8, int16, int32, int64
	 */
	//int8 es para almacenar numeros enteros pequenos
	var maximoInt8 int8 = math.MaxInt8 // maximo
	var minimoInt8 int8 = math.MinInt8 //minimo
	// uint: solo almacena datos numericos postivos
	var uintValoresPositivos uint = 10
	fmt.Println("Maximo int8: ", maximoInt8, ", Minimo int8: ", minimoInt8)
	fmt.Println("uint es solo para valores positivos", uintValoresPositivos)

	fmt.Println("valor minimo int64", math.MaxInt64, " y valor maximo de int 64", math.MaxInt64)

	/*
	*Datos numeros decimal
	* float32 y float64
	* float32 = para numeros pequenos
	* float64 = para numeros grandes
	 */

	var f32 float32 = math.MaxFloat32
	var float64 float64 = math.MaxFloat64
	fmt.Println("Maximo valor float32 ", f32)
	fmt.Println("Maximo valor float64 ", float64)

	//Booleanos
	//var valueBool bool = true

	//Cadena de caracteres
	fullName := "Lautaro Gallo \t(Alias: \"lautarovg02\")\n"
	fmt.Println(fullName)

	/*
	 *Tipo byte
	 *almacena numeros positivos hata 255
	 *utilizado normalmente para guardar datos de caracteres ascii
	 */
	var a byte = 'a'
	// no va a imprimir a, sino que va a imprimir el valro de codigo asci, seria 97
	fmt.Println(a)

	//acceder a valores de una cadena
	vocales := "aeiou"
	fmt.Println(vocales[0], "= imprime el valor b")

	/*
	*Tipo rune
	* almacena numeros con un tamano de int32
	* Representa un punto de codigo Unicode en GO
	 */

	var r rune = '💘'
	fmt.Println("Corazon valor unicode, ", r)

	//Valores predertiminados o valor cero

	var (
		defaultInt     int
		defaultBool    bool
		defaultFloat32 float32
		defaultUint    uint
		defaultString  string
	)
	fmt.Println("Valores por defercto ", defaultInt, defaultUint, defaultFloat32, defaultString, defaultBool)
}
