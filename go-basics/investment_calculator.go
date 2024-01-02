package main

import (
	"fmt"
	"math"
)

func main() {
	const inflationRate float64 = 2.5
	// var investmentAmount float64 = 1000
	// expectedReturnRate := 5.5
	// years := 10.0
	// inflationRate = 4.5 // This won't work

	var investmentAmount float64
	var years float64
	var expectedReturnRate float64

	fmt.Print("Enter the value for Investment Amount: ")
	fmt.Scan(&investmentAmount)

	fmt.Print("Enter the value for Expected Return Rate: ")
	fmt.Scan(&expectedReturnRate)

	fmt.Print("Enter the years: ")
	fmt.Scan(&years)

	futureValue := investmentAmount * math.Pow((1+expectedReturnRate/100), float64(years))
	futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)

	// fmt.Println("Enter Future Value: ", futureValue)
	// fmt.Println("Inflated Value (adjusted for inflation):", futureRealValue)
	// fmt.Printf("Future value: %v \nFuture value (adjusted for inflation): %v\n", futureValue, futureRealValue)
	// fmt.Printf("Future value: %.0f \nFuture value (adjusted for inflation): %.1f\n", futureValue, futureRealValue)

	// ! Creating Formatted Strings
	formattedFV := fmt.Sprintf("Future value: %.1f\n", futureValue)
	formattedRFV := fmt.Sprintf("Future value (adjusted for inflation): %.2f\n", futureRealValue)

	fmt.Print(formattedFV)
	fmt.Print(formattedRFV)

	// ! Multiline Formatted Strings
	fmt.Printf(`Future Value: %.1f
Future Value (adjusted for inflation): %.1f`, futureValue, futureRealValue)

}
