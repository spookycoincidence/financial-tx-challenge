package main

import (
	"fmt"
	"net/smtp"
)

func enviarCorreo(resumen Resumen) error {
	emailBody := generarCuerpoCorreo(resumen)

	d := smtp.PlainAuth("", "correofalso123@gmail.com", "tucontraseña", "smtp.gmail.com")

	err := smtp.SendMail("smtp.gmail.com:587", d, "correofalso123@gmail.com", []string{"destinatariofalso@gmail.com"}, []byte(emailBody))
	if err != nil {
		return err
	}

	fmt.Println("Correo electrónico enviado con éxito.")
	return nil
}

func generarCuerpoCorreo(resumen Resumen) string {
	emailBody := `
		<!DOCTYPE html>
		<html>
		<head>
			<style>
				body {
					font-family: Arial, sans-serif;
					background-color: #f4f4f4;
					padding: 20px;
				}
				.container {
					background-color: #fff;
					border-radius: 5px;
					padding: 20px;
					box-shadow: 0 0 10px rgba(0, 0, 0, 0.1);
				}
				h1 {
					color: #333;
				}
				.section {
					margin-bottom: 20px;
				}
				.table {
					width: 100%;
					border-collapse: collapse;
				}
				.table th, .table td {
					padding: 10px;
					border: 1px solid #ddd;
					text-align: left;
				}
			</style>
		</head>
		<body>
			<div class="container">
				<h1>Resumen de Transacciones</h1>
				<div class="section">
					<h2>Total Balance</h2>
					<p>$%.2f</p>
				</div>
				<div class="section">
					<h2>Número de Transacciones por Mes</h2>
					<table class="table">
						<tr>
							<th>Mes</th>
							<th>Número de Transacciones</th>
						</tr>
						%s
					</table>
				</div>
				<div class="section">
					<h2>Promedio de Montos de Débito por Mes</h2>
					<table class="table">
						<tr>
							<th>Mes</th>
							<th>Promedio de Débito</th>
						</tr>
						%s
					</table>
				</div>
				<div class="section">
					<h2>Promedio de Montos de Crédito por Mes</h2>
					<table class="table">
						<tr>
							<th>Mes</th>
							<th>Promedio de Crédito</th>
						</tr>
						%s
					</table>
				</div>
				<img src="https://stori.com/stori_logo.png" alt="Stori Logo" width="200" height="100">
			</div>
		</body>
		</html>
	`

	numTransacciones := ""
	for month, num := range resumen.NumTransaccionesPorMes {
		numTransacciones += fmt.Sprintf("<tr><td>%s</td><td>%d</td></tr>", month, num)
	}

	promedioDebito := ""
	for month, avgDebit := range resumen.PromedioDebitoPorMes {
		promedioDebito += fmt.Sprintf("<tr><td>%s</td><td>$%.2f</td></tr>", month, avgDebit)
	}

	promedioCredito := ""
	for month, avgCredit := range resumen.PromedioCreditoPorMes {
		promedioCredito += fmt.Sprintf("<tr><td>%s</td><td>$%.2f</td></tr>", month, avgCredit)
	}

	emailBody = fmt.Sprintf(emailBody, resumen.SaldoTotal, numTransacciones, promedioDebito, promedioCredito)

	return emailBody
}
