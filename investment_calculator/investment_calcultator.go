package main

import (
	"fmt"
	"math"
)

const inflationRate = 2.5

func main() {
	var investmentAmount float64

	var expectedReturnRate = 5.5
	var years float64

	ouputText("Insert investment Amount: ")
	fmt.Scan(&investmentAmount)

	ouputText("Expected return rate: ")
	fmt.Scan(&expectedReturnRate)

	ouputText("Years: ")
	fmt.Scan(&years)

	futureValue, futureRealValue := calculateFutureValues(investmentAmount, expectedReturnRate, years)
	// futureValue := investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	// futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)

	fmt.Println(futureValue)
	fmt.Println(futureRealValue)
}

func ouputText(text string) {
	fmt.Print(text)
}

func calculateFutureValues(investmentAmount, expectedReturnRate, years float64) (fv float64, frv float64) {
	fv = investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	frv = fv / math.Pow(1+inflationRate/100, years)
	return fv, frv
	// return
}
