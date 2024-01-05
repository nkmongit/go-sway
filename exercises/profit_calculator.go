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

	// fmt.Print("Enter the revenue: ")
	// fmt.Scan(&revenue)
	revenue, err := getUserInput("Enter the revenue: ")
	if err != nil {
		fmt.Println(err)
		return
	}

	// fmt.Print("Enter the expenses: ")
	// fmt.Scan(&expenses)
	expenses, err := getUserInput("Enter the expenses: ")
	if err != nil {
		fmt.Println(err)
		return
	}
	// fmt.Print("Enter the tax rate: ")
	// fmt.Scan(&taxRate)
	taxRate, err := getUserInput("Enter the tax-rate: ")
	if err != nil {
		fmt.Println(err)
		return
	}
	// revenue, expenses, taxRate = takeInputs(revenue, expenses, taxRate)

	// ebt := revenue - expenses
	// profit := ebt * (1 - taxRate/100)
	// ratio := ebt / profit

	ebt, profit, ratio := calculateThings(revenue, expenses, taxRate)

	fmt.Printf("Your EBT: %.1f \n", ebt)
	fmt.Printf("Your profit: %.1f \n", profit)
	fmt.Printf("Your ratio: %.2f \n", ratio)
	storeResults(ebt, profit, ratio)
}

func storeResults(ebt, profit, ratio float64) {
	results := fmt.Sprintf("EBT: %.1f\nProfit: %.1f\nRatio: %.3f\n", ebt, profit, ratio)
	os.WriteFile("results.txt", []byte(results), 0644)
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

func getUserInput(infoText string) (float64, error) {
	var userInput float64
	fmt.Print(infoText)
	fmt.Scan(&userInput)

	if userInput < 1 {
		return 0, errors.New("Invalid Input")
	}
	return userInput, nil
}
