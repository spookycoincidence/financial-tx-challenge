package main

import (
	"log"
	"strconv"
)

// atoi convierte una cadena a un entero
func atoi(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		log.Fatal(err)
	}
	return i
}
