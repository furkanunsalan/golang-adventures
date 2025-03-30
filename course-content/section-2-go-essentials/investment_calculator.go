package main

import (
	"fmt"
	"math"
)

const inflationRate = 2.5

func main() {
	var investmentAmout float64
	var years float64
	var expectedReturnRate float64
	// var years float64 = 10

	// fmt.Print("Investment Amout: ")
	outputText("Investment Amout: ")
	fmt.Scan(&investmentAmout)

	// fmt.Print("Expected Return Rate: ")
	outputText("Expected Return Rate: ")
	fmt.Scan(&expectedReturnRate)

	// fmt.Print("Years of investment: ")
	outputText("Years of investment: ")
	fmt.Scan(&years)

	// futureValue := investmentAmout * math.Pow((1+expectedReturnRate/100), years)
	// futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)

	futureValue, futureRealValue := calculateFutureValues(investmentAmout, expectedReturnRate, years)

	// fmt.Println("future value of your money: ", futureValue)
	// fmt.Println("Future valuation with inflation: ", futureRealValue)
	// fmt.Printf("Future Value: %.2f\nFuture valuation with inflation: %.2f", futureValue, futureRealValue)

	formattedString := fmt.Sprintf("Future Value: %.2f\nFuture valuation with inflation: %.2f", futureValue, futureRealValue)

	// fmt.Printf(`Future Value: %.2f
	// Future valuation with inflation: %.2f`, futureValue, futureRealValue)

	fmt.Println(formattedString)
}

func outputText(text string) {
	fmt.Print(text)
}

func calculateFutureValues(investmentAmout float64, expectedReturnRate float64, years float64) (float64, float64) {
	fv := investmentAmout * math.Pow((1+expectedReturnRate/100), years)
	rfv := fv / math.Pow(1+inflationRate/100, years)
	return fv, rfv
}

// Alternative return syntax
// remove ":" from the variable declarations and rather
// declare them inside the return parantheses

// func calculateFutureValues(investmentAmout float64, expectedReturnRate float64, years float64) (fv float64, rfv float64) {
// 	fv = investmentAmout * math.Pow((1+expectedReturnRate/100), years)
// 	rfv = fv / math.Pow(1+inflationRate/100, years)
// 	return fv, rfv
// }
