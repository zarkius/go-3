package main

import (
	"fmt"
	"net/http"
	"sync"
)

// Variable global para almacenar el estado del semáforo y un mutex para proteger su acceso
var (
	estadoSemaforo int = 0 // 0 = cerrado, 1 = abierto
	mutex          sync.Mutex
)

// Manejador del endpoint del semáforo
func semaforoHandler(w http.ResponseWriter, r *http.Request) {
	// Solo permitir el cambio de estado a través de una solicitud GET para simplificar
	if r.Method == "GET" {
		if estado, ok := r.URL.Query()["estado"]; ok {
			// Validar y actualizar el estado del semáforo de manera segura
			mutex.Lock()
			if estado[0] == "1" {
				estadoSemaforo = 1
			} else if estado[0] == "0" {
				estadoSemaforo = 0
			}
			mutex.Unlock()
		}
	} else {
		// Método HTTP no permitido
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	// Devolver el estado actual del semáforo
	mutex.Lock()
	fmt.Fprintf(w, "Estado del semáforo: %d", estadoSemaforo)
	mutex.Unlock()
}

func main() {
	// Asociar el manejador al endpoint "/semaforo"
	http.HandleFunc("/semaforo", semaforoHandler)

	// Iniciar el servidor en el puerto 8080
	fmt.Println("Servidor iniciado en http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
