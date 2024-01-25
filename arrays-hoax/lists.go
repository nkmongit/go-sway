package main

import "fmt"

func learnSlices() {
	prices := []float64{10.6, 8.0, 8.7}
	fmt.Println(prices)
	prices = append(prices, 12.0)

	fmt.Println(prices)
}

func main() {
	fmt.Println("ARRAYS")

	var productNames [4]string = [4]string{"A Book"}
	prices := [4]float64{10.99, 9.99, 45.99, 20.0}

	productNames[2] = "A Carpet"

	fmt.Println(prices)
	fmt.Println(productNames)

	fmt.Println(prices[0])

	// SLICING ARRAYS
	// featuredPrices := prices[1:3]

	// featuredPrices := prices[:3]
	featuredPrices := prices[1:]
	// featuredPrices = prices[:]

	featuredPrices[0] = 7
	fmt.Println(featuredPrices)
	fmt.Println(prices)

	// using len() and cap()
	fmt.Println(len(featuredPrices), cap(featuredPrices))

	fmt.Println("----------------DYNAMIC SLICES-----------------")
	learnSlices()
}
