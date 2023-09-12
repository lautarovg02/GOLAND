package main

import (
	"fmt"
	"math/rand"
)

func main() {
	play()
}

func playAgain() {
	var eleccion string
	fmt.Println("\nDesea jugar denuevo? Coloque si o no: ")
	fmt.Scanln(&eleccion)
	switch eleccion {
	case "si":
		play()
	case "no":
		fmt.Println("Okey, muchas gracias por jugar, nos vemos")
	default:
		fmt.Println("Opcion invalida: ingree nuevamante")
		playAgain()
	}
}

func play() {
	numRandom := rand.Intn(100)
	var numEntered int
	var numAttempts int
	const MAXATTEMPTS = 10

	for numAttempts < MAXATTEMPTS {
		fmt.Printf("Ingrese un numero e intente adivinar un numero random (intentos restantes: %d): \n", MAXATTEMPTS-numAttempts)
		fmt.Scanln(&numEntered)
		numAttempts++
		if numEntered == numRandom {
			fmt.Printf("Felicitaciones ah ganado, logro adivinar el numero en %d \n", MAXATTEMPTS-numAttempts)
			playAgain()
		} else if numEntered > numRandom {
			fmt.Printf("Pista: %d  es mayor al numero a adivinar\n", numEntered)
		} else if numEntered < numRandom {
			fmt.Printf("Pista: %d es menor al numero a adivinar\n", numEntered)
		}
	}
	fmt.Printf("Se acabaron los intentos el numero era: %d", numRandom)
	playAgain()
}
