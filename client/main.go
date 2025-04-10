package main

import (
	"client/globals"
	"client/utils"
	"log"
)

func main() {
	utils.ConfigurarLogger()

	// loggear "Hola soy un log" usando la biblioteca log
	log.Println("Hola soy un log")

	globals.ClientConfig = utils.IniciarConfiguracion("config.json")

	// validar que la config este cargada correctamente
	if globals.ClientConfig == nil {
		log.Println("Error: La configuración no se cargó correctamente.")
		return
	}

	// loggeamos el valor de la config
	log.Println("Valor de la configuración:", globals.ClientConfig)

	// ADVERTENCIA: Antes de continuar, tenemos que asegurarnos que el servidor esté corriendo para poder conectarnos a él

	// enviar un mensaje al servidor con el valor de la config       (??????????)

	// leer de la consola el mensaje

	utils.LeerConsola()
	log.Println("Mensaje leído de la consola")

	// generamos un paquete y lo enviamos al servidor
	utils.GenerarYEnviarPaquete()
	log.Println("Paquete enviado al servidor")
}
