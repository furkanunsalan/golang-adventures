package main

import (
	"fmt"
	"math"
)

func main() {
	const inflationRate = 2.5
	var investmentAmout float64
	var years float64
	var expectedReturnRate float64
	// var years float64 = 10

	fmt.Print("Investment Amout: ")
	fmt.Scan(&investmentAmout)

	fmt.Print("Expected Return Rate: ")
	fmt.Scan(&expectedReturnRate)

	fmt.Print("Years of investment: ")
	fmt.Scan(&years)


	futureValue := investmentAmout * math.Pow((1+expectedReturnRate/100), years)
	futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)

	fmt.Println("future value of your money: ", futureValue)
	fmt.Println("future validation with inflation: ", futureRealValue)
}
