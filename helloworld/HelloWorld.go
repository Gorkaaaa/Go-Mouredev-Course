// Primer comando (crear proyecto): go mod init helloworld

package main // Cargar el paquete principal.

import "fmt" // Modulo

// Funcion principal de la página.
func main() {
	print("Hello world")     // Print método normal.
	fmt.Println("Hola, Go.") // Print utilizando la libreria.
}
