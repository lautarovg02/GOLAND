package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {

	// nos devuelve en que sistema operativo se esta ejecutando go
	// A lo mismo que con if, tmb se pueden inicializar  variables y utilizarlas ahi

	switch runtime.GOOS {
	case "windows":
		fmt.Println("Go run -> windows")
	case "linux":
		fmt.Println("Go run -> linux")
	case "darwin":
		fmt.Println("Go run -> macOs")
	default:
		fmt.Println("Go run -> otro OS")
	}

	// O asi

	//os := runtime.GOOS
	//
	//switch os {
	//case "windows":
	//	fmt.Println("Go run -> windows")
	//case "linux":
	//	fmt.Println("Go run -> linux")
	//case "darwin":
	//	fmt.Println("Go run -> macOs")
	//default:
	//	fmt.Println("Go run -> otro OS")
	//}

	switch t := time.Now(); {
	case t.Hour() < 12:
		fmt.Println("Buen dia, son las ", t.Hour())
	case t.Hour() < 17:
		fmt.Println("Buenas tardes, son las ", t.Hour())
	default:
		fmt.Println("Buenas noches, son las ", t.Hour())
	}
}
