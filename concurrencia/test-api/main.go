package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	start := time.Now()
	apis := []string{
		"https://management.azure.com",
		"https://dev.azure.com",
		"https://outlook.office.com/",
		"https://api.somewhereintheinternet.com",
		"https://graph.microsoft.com",
	}
	for _, api := range apis {
		go checkAPI(api) //Aqui se aplica la concurrencia agregagando go
	}
	time.Sleep(5 * time.Second) //Se agrega un tiempo de suspenso para definir el tiempo que va esperar  para la ejecucion
	elapsed := time.Since(start)
	fmt.Printf("¡Listo! ¡Tomo %v segundos!\n", elapsed.Seconds())

}

func checkAPI(api string) {
	if _, err := http.Get(api); err != nil {
		fmt.Printf("ERROR: !%s esta caido\n", api)
	} else {
		fmt.Printf("OK: %s esta funcionando\n", api)
	}
}
