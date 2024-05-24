package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

type Transaction struct {
	ID     int
	Date   time.Time
	Amount float64
}

func main() {
	file, err := os.Open("txns.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	reader := csv.NewReader(file)

	lines, err := reader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	// Mapa para almacenar transacciones agrupadas por mes
	transactionsByMonth := make(map[string][]Transaction)

	var totalBalance float64
	transactionCountByMonth := make(map[string]int)
	totalDebitByMonth := make(map[string]float64)
	totalCreditByMonth := make(map[string]float64)

	currentYear := time.Now().Year()

	for _, line := range lines {
		if line[0] == "Id" {
			continue
		}

		id, _ := strconv.Atoi(line[0])
		dateString := fmt.Sprintf("%d-%s", currentYear, line[1])
		date, _ := time.Parse("2006-1/2", dateString)
		amount, _ := strconv.ParseFloat(strings.TrimPrefix(line[2], "+-"), 64)

		if strings.HasPrefix(line[2], "+") {
			totalCreditByMonth[date.Format("2006-01-02")] += amount
		} else {
			totalDebitByMonth[date.Format("2006-01-02")] += amount
		}

		// Actualizar el saldo total de la cuenta
		totalBalance += amount

		transactionCountByMonth[date.Format("2006-01-02")]++

		txn := Transaction{
			ID:     id,
			Date:   date,
			Amount: amount,
		}

		transactionsByMonth[date.Format("2006-01-02")] = append(transactionsByMonth[date.Format("2006-01-02")], txn)
	}

	fmt.Printf("Total balance is %.2f\n", totalBalance)

	for month, count := range transactionCountByMonth {
		fmt.Printf("Number of transactions in %s: %d\n", month, count)
	}

	for month, totalDebit := range totalDebitByMonth {
		averageDebit := totalDebit / float64(transactionCountByMonth[month])
		fmt.Printf("Average debit amount in %s: %.2f\n", month, averageDebit)
	}

	for month, totalCredit := range totalCreditByMonth {
		averageCredit := totalCredit / float64(transactionCountByMonth[month])
		fmt.Printf("Average credit amount in %s: %.2f\n", month, averageCredit)
	}
}
