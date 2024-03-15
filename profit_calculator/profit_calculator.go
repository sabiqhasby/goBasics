package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {
	// var revenue float64
	// var expenses float64
	// var taxRate float64

	revenue, err := getUserInput("Revenue: ")
	if err != nil {
		panic(err)
	}
	expenses, err := getUserInput("Expenses: ")
	if err != nil {
		panic(err)
	}
	taxRate, err := getUserInput("Tax Rate: ")
	if err != nil {
		panic(err)
	}

	earningsBeforeTax, profit, ratio := calculateFinancial(revenue, expenses, taxRate)
	// earningsBeforeTax := revenue - expenses
	// profit := earningsBeforeTax * (1 - taxRate/100)

	// ratio := earningsBeforeTax / profit

	fmt.Printf("ebt: %.1f\n", earningsBeforeTax)
	fmt.Printf("profit: %.1f\n", profit)
	fmt.Printf("ratio: %.1f\n", ratio)

	storeResults(earningsBeforeTax, profit, ratio)
}

func calculateFinancial(revenue, expenses, taxRate float64) (float64, float64, float64) {
	earningsBeforeTax := revenue - expenses
	profit := earningsBeforeTax * (1 - taxRate/100)

	ratio := earningsBeforeTax / profit

	return earningsBeforeTax, profit, ratio
}

func getUserInput(text string) (float64, error) {
	var userInput float64
	fmt.Print(text)
	fmt.Scan(&userInput)

	if userInput <= 0 {
		return 0, errors.New("cant use negative or zero value")
	}

	return userInput, nil
}

func storeResults(earningsBeforeTax, profit, ratio float64) {
	results := fmt.Sprintf("EBT: %.1f\nProfit: %.1f\nRatio: %.3f\n", earningsBeforeTax, profit, ratio)
	os.WriteFile("results.txt", []byte(results), 0644)
}
