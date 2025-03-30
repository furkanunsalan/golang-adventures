package main

import "fmt"

func main() {
	// var revenue float64
	// var expenses float64
	// var taxRate float64

	revenue := getUserInput("Revenue: ")
	expenses := getUserInput("Expenses: ")
	taxRate := getUserInput("Tax Rate: ")

	ebt, profit, ratio := calculateFinancials(revenue, expenses, taxRate)

	fmt.Printf("Before Tax: %.2f\n", ebt)
	fmt.Printf("After Tax (Profit Amount): %.2f\n", profit)
	fmt.Printf("Profit Ratio: %.2f\n", ratio)
}

func calculateFinancials(revenue, expenses, taxRate float64) (float64, float64, float64) {
	ebt := revenue - expenses
	profit := ebt * (1 - taxRate/100)
	ratio := ebt / profit

	return ebt, profit, ratio
}

func getUserInput(infoText string) float64 {
	var userInput float64
	fmt.Print(infoText)
	fmt.Scan(&userInput)

	return userInput
}
