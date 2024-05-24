package main

import "fmt"

// Aca simulamos ficticiamente el almacenamiento del resumen de la transaccion en AWS
func AlmacenarResumenEnS3(resumen Resumen) {
	fmt.Println("Los datos del resumen se han almacenado correctamente en Amazon S3.")
}
