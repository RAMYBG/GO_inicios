package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	t := time.Now()
	hora := t.Hour()
	fmt.Println(t.Hour())
	if hora < 12 {
		fmt.Println("!Mañana¡")
	} else if hora < 17 {
		fmt.Println("!Tarde¡")
	} else {
		fmt.Println("!Tarde¡")
	}

	//Utilizar switch
	os := runtime.GOOS
	switch os {
	case "windows":
		fmt.Println("Go run -> Windows")
	case "linux":
		fmt.Println("Go run -> Linux")
	case "darwin":
		fmt.Println("Go run -> macOS")
	default:
		fmt.Println("Go run -> Otro OS")
	}
}
