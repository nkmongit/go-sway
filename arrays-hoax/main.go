package main

import "fmt"

// creating a type
type floatMap map[string]float64

func (m floatMap) output() {
	fmt.Println(m)
}

func main() {
	userNames := []string{}
	userNames = append(userNames, "Nishant")
	userNames = append(userNames, "Mohapatra")

	firstNames := make([]string, 2, 2)

	firstNames = append(firstNames, "Dishant")
	firstNames = append(firstNames, "Dishant")

	fmt.Println(firstNames)
	fmt.Println(userNames)

	// creating maps using `make`

	// courseRatings := make(map[string]float64, 3)

	courseRatings := make(floatMap, 5)

	courseRatings["go"] = 4.7
	courseRatings["react"] = 4.8

	fmt.Println(courseRatings)

	courseRatings.output()

	// loops for maps, slices and array
	for i, v := range userNames {
		fmt.Println("Index: ", i)
		fmt.Println("Value: ", v)
	}

	// maps
	for key, value := range courseRatings {
		fmt.Println(key, ":", value)
	}
}
