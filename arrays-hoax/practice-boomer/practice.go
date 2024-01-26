package main

import (
	"fmt"
)

func main() {
	// 1) Create a new array (!) that contains three hobbies you have
	// 		Output (print) that array in the command line.
	hobbies := [3]string{"coding", "sleeping", "yawning"}
	fmt.Println(hobbies)

	// 2) Also output more data about that array:
	//		- The first element (standalone)
	//		- The second and third element combined as a new list
	fmt.Println(hobbies[0])
	fmt.Println(hobbies[1:])

	// 3) Create a slice based on the first element that contains
	//		the first and second elements.
	//		Create that slice in two different ways (i.e. create two slices in the end)

	sliced1 := hobbies[0:3]
	sliced2 := hobbies[:3]

	fmt.Println(sliced1)
	fmt.Println(sliced2)

	// 4) Re-slice the slice from (3) and change it to contain the second
	//		and last element of the original array.

	resliced := sliced1[1:3]
	fmt.Println(resliced)
	// 5) Create a "dynamic array" that contains your course goals (at least 2 goals)
	goals := []string{"deterministic", "accountable"}
	// 6) Set the second goal to a different one AND then add a third goal to that existing dynamic array
	goals[1] = "learner"
	goals = append(goals, "confident")

	fmt.Println(goals)

	// 7) Bonus: Create a "Product" struct with title, id, price and create a
	//		dynamic list of products (at least 2 products).
	//		Then add a third product to the existing list of products.

	type Product struct {
		title string
		id    int
		price float64
	}

	pro := []Product{
		{
			title: "Cookies",
			id:    1,
			price: 3.9,
		},
		{
			title: "Baking Powder",
			id:    2,
			price: 20.8,
		},
	}

	pro = append(pro, Product{
		title: "Shampoo",
		id:    3,
		price: 45,
	})
	fmt.Println(pro)
}
