package main

import (
	"fmt"
	"time"
)

type TransactionDB struct {
	Transactions []Transaction
}

func (db *TransactionDB) ConsultarTransaccionesPorRangoFecha(desde, hasta time.Time) []Transaction {
	var resultado []Transaction
	for _, txn := range db.Transactions {
		if txn.Date.After(desde) && txn.Date.Before(hasta) {
			resultado = append(resultado, txn)
		}
	}
	return resultado
}

func (db *TransactionDB) InsertarTransaccion(txn Transaction) {
	db.Transactions = append(db.Transactions, txn)
}

func (db *TransactionDB) ActualizarTransaccion(id int, nuevaTxn Transaction) error {
	for i, txn := range db.Transactions {
		if txn.ID == id {
			db.Transactions[i] = nuevaTxn
			return nil
		}
	}
	return fmt.Errorf("transacción con ID %d no encontrada", id)
}

func (db *TransactionDB) EliminarTransaccion(id int) error {
	for i, txn := range db.Transactions {
		if txn.ID == id {
			// Eliminar la transacción del slice
			db.Transactions = append(db.Transactions[:i], db.Transactions[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("transacción con ID %d no encontrada", id)
}
