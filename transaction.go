package main

import (
	"strconv"
	"time"
)

type Transaction struct {
	ID     int
	Date   time.Time
	Amount float64
}

func atoi(s string) (int, error) {
	i, err := strconv.Atoi(s)
	if err != nil {
		return 0, err
	}
	return i, nil
}

func generarResumen(transacciones map[string][]Transaction) Resumen {
	resumen := Resumen{}

	resumen.SaldoTotal = calcularSaldoTotal(transacciones)
	resumen.NumTransaccionesPorMes = calcularNumTransaccionesPorMes(transacciones)
	resumen.PromedioDebitoPorMes, resumen.PromedioCreditoPorMes = calcularPromedioMontosPorMes(transacciones)

	return resumen
}

func calcularSaldoTotal(transacciones map[string][]Transaction) float64 {
	total := 0.0
	for _, txns := range transacciones {
		for _, txn := range txns {
			total += txn.Amount
		}
	}
	return total
}

func calcularNumTransaccionesPorMes(transacciones map[string][]Transaction) map[string]int {
	numTransacciones := make(map[string]int)
	for month, txns := range transacciones {
		numTransacciones[month] = len(txns)
	}
	return numTransacciones
}

func calcularPromedioMontosPorMes(transacciones map[string][]Transaction) (map[string]float64, map[string]float64) {
	promedioDebito := make(map[string]float64)
	promedioCredito := make(map[string]float64)
	for month, txns := range transacciones {
		var totalDebito, totalCredito float64
		numDebito, numCredito := 0, 0
		for _, txn := range txns {
			if txn.Amount < 0 {
				totalDebito += txn.Amount
				numDebito++
			} else {
				totalCredito += txn.Amount
				numCredito++
			}
		}
		if numDebito > 0 {
			promedioDebito[month] = totalDebito / float64(numDebito)
		} else {
			promedioDebito[month] = 0
		}
		if numCredito > 0 {
			promedioCredito[month] = totalCredito / float64(numCredito)
		} else {
			promedioCredito[month] = 0
		}
	}
	return promedioDebito, promedioCredito
}
