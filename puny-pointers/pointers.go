package main

import "fmt"

func main() {
	age := 32 // regular variable
	fmt.Println("Age:", age)

	agePointer := &age

	adultYears := getAdultYears(age)
	fmt.Println(adultYears)

	// De-referencing the pointer here
	fmt.Println("Age:", *agePointer)
	// using pointer
	adultYearsPointer := getAdultYearsPointer(agePointer)
	fmt.Println("Age using pointer:", adultYearsPointer)
}

// age passed in the params
// will have different address as of the value age in main function
// because a copy has created.
func getAdultYears(age int) int {
	return age - 18
}

func getAdultYearsPointer(agePointer *int) int {
	return (*agePointer - 18)
}
