package main

import "fmt"

func main() {
	// var revenue float64
	// var expenses float64
	// var taxRate float64

	// fmt.Print("Enter the revenue: ")
	// fmt.Scan(&revenue)
	revenue := getUserInput("Enter the revenue: ")

	// fmt.Print("Enter the expenses: ")
	// fmt.Scan(&expenses)
	expenses := getUserInput("Enter the expenses: ")

	// fmt.Print("Enter the tax rate: ")
	// fmt.Scan(&taxRate)
	taxRate := getUserInput("Enter the tax-rate: ")

	// revenue, expenses, taxRate = takeInputs(revenue, expenses, taxRate)

	// ebt := revenue - expenses
	// profit := ebt * (1 - taxRate/100)
	// ratio := ebt / profit

	ebt, profit, ratio := calculateThings(revenue, expenses, taxRate)

	fmt.Print("Your EBT: %.1f \n", ebt)
	fmt.Print("Your profit: %.1f \n", profit)
	fmt.Print("Your ratio: %.2f \n", ratio)
}

func calculateThings(revenue, expenses, taxRate float64) (float64, float64, float64) {
	ebt := revenue - expenses
	profit := ebt * (1 - taxRate/100)
	ratio := ebt / profit
	return ebt, profit, ratio
}

// func takeInputs(revenue, expenses, taxRate float64) (float64, float64, float64) {
// 	fmt.Print("Enter the value for revenue: ")
// 	fmt.Scan(&revenue)

// 	fmt.Print("Enter the value for expenses: ")
// 	fmt.Scan(&expenses)

// 	fmt.Print("Enter the value for tax-rate: ")
// 	fmt.Scan(&taxRate)

// 	return revenue, expenses, taxRate
// }

func getUserInput(infoText string) float64 {
	var userInput float64
	fmt.Print(infoText)
	fmt.Scan(&userInput)
	return userInput
}
