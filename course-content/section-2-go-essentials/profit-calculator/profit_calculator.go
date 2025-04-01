package main

import (
	"errors"
	"fmt"
	"os"
)

// GOALS
// 1) Validate user input
// 		=> Show error message & exit if invalid input is provided
//		- Not negative numbers
//		- Not 0
// 2) Store calculated results into file

func main() {
	// var revenue float64
	// var expenses float64
	// var taxRate float64

	revenue, err1 := getUserInput("Revenue: ")

	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

	expenses, err2 := getUserInput("Expenses: ")

	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

	taxRate, err3 := getUserInput("Tax Rate: ")

	if err1 != nil || err2 != nil || err3 != nil {
		fmt.Println(err1)
		return
	}

	ebt, profit, ratio := calculateFinancials(revenue, expenses, taxRate)

	fmt.Printf("Before Tax: %.2f\n", ebt)
	fmt.Printf("After Tax (Profit Amount): %.2f\n", profit)
	fmt.Printf("Profit Ratio: %.2f\n", ratio)
	storeResults(ebt, profit, ratio)
}

func storeResults(ebt, profit, ratio float64) {
	results := fmt.Sprintf("EBT: %.1f\nProfit: %.1f\nRatio: %.3f\n", ebt, profit, ratio)
	os.WriteFile("results.txt", []byte(results), 0644)
}

func calculateFinancials(revenue, expenses, taxRate float64) (float64, float64, float64) {
	ebt := revenue - expenses
	profit := ebt * (1 - taxRate/100)
	ratio := ebt / profit

	return ebt, profit, ratio
}

func getUserInput(infoText string) (float64, error) {
	var userInput float64
	fmt.Print(infoText)
	fmt.Scan(&userInput)

	if userInput <= 0 {
		return 0, errors.New("value must be positive number")
	}

	return userInput, nil
}
