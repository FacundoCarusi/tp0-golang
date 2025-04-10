package main

import (
	"log"
	"net/http"
	"server/utils"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/paquetes", utils.RecibirPaquetes)
	mux.HandleFunc("/mensaje", utils.RecibirMensaje)

	/*panic("no implementado!")
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		panic(err)
	}*/

	http.ListenAndServe(":8080", mux) // Inicia el servidor HTTP en el puerto 8080
	log.Println("Servidor escuchando en el puerto 8080")
}
