package main // Define el paquete principal de este programa en Go.

import ( // Importa múltiples paquetes
	"fmt"      // Importa el paquete fmt para la entrada y salida de datos
	"net/http" // Importa el paquete net/http para realizar solicitudes HTTP
	"time"     // Importa el paquete time para trabajar con fechas y horas
)

func main() { // Define la función principal `main`, que es el punto de entrada de cualquier programa en Go.
	start := time.Now() // Registra el tiempo de inicio de la ejecución del programa

	apis := []string{ // Define un slice de strings con URLs de diferentes APIs
		"https://management.azure.com",
		"https://dev.azure.com",
		"https://outlook.office.com/",
		"https://api.somewhereintheinternet.com",
		"https://graph.microsoft.com",
	}

	ch := make(chan string) // Crea un canal de tipo string para comunicar información entre goroutines

	for _, api := range apis { // Itera sobre cada API en el slice de apis
		go checkAPI(api, ch) // Lanza una goroutine para cada API, ejecutando la función checkAPI de manera concurrente
	}

	for i := 0; i < len(apis); i++ { // Itera sobre el número de APIs para recibir mensajes del canal
		fmt.Print(<-ch) // Imprime los mensajes recibidos a través del canal
	}

	elapsed := time.Since(start)                                  // Calcula el tiempo transcurrido desde el inicio de la ejecución
	fmt.Printf("¡Listo! ¡Tomo %v segundos!\n", elapsed.Seconds()) // Imprime el tiempo total tomado para verificar todas las APIs

}

func checkAPI(api string, ch chan string) { // Define la función checkAPI que recibe una URL de API y un canal
	if _, err := http.Get(api); err != nil { // Intenta realizar una solicitud GET a la API y verifica si hay errores
		ch <- fmt.Sprintf("ERROR: !%s esta caido\n", api) // Envía un mensaje de error al canal si la API está caída
	} else {
		ch <- fmt.Sprintf("OK: %s esta funcionando\n", api) // Envía un mensaje de éxito al canal si la API está funcionando
	}
}
