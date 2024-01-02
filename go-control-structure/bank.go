package main

import "fmt"

func main() {

	var accountBalance = 1000.0

	fmt.Println("Welcome to Go Bank!")
	for {
		fmt.Println("What do you want to do?")

		fmt.Println("1. Check balance")
		fmt.Println("2. Deposit money")
		fmt.Println("3. Withdraw money")
		fmt.Println("4. Exit")

		var choice int
		fmt.Print("Your choice: ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			fmt.Println("Your balance is", accountBalance)
		case 2:
			fmt.Print("How much money you want to deposit? ")
			var depositAmount float64
			fmt.Scan(&depositAmount)

			if depositAmount <= 0 {
				fmt.Println("Invalid amount. Must be graeter than 0")
				continue
			}

			accountBalance += depositAmount
			fmt.Println("Balance updated! New amount:", accountBalance)
		case 3:
			fmt.Print("How much money your want to withdraw? ")
			var withdrawAmount float64
			fmt.Scan(&withdrawAmount)

			if withdrawAmount <= 0 {
				fmt.Println("Invalid amount. Must be graeter than 0")
				continue
			}

			if withdrawAmount > accountBalance {
				fmt.Println("You don't have enough balance to withdraw")
				continue
			}

			accountBalance -= withdrawAmount
			fmt.Println("Balance updated! New amount:", accountBalance)
		default:
			fmt.Println("Goodbye!")
			fmt.Println("Thanks for choosing our bank!")
			return
		}

		// if choice == 1 {
		// 	fmt.Println("Your balance is", accountBalance)
		// } else if choice == 2 {
		// 	fmt.Print("How much money you want to deposit? ")
		// 	var depositAmount float64
		// 	fmt.Scan(&depositAmount)

		// 	if depositAmount <= 0 {
		// 		fmt.Println("Invalid amount. Must be graeter than 0")
		// 		continue
		// 	}

		// 	accountBalance += depositAmount
		// 	fmt.Println("Balance updated! New amount:", accountBalance)
		// } else if choice == 3 {
		// 	fmt.Print("How much money your want to withdraw? ")
		// 	var withdrawAmount float64
		// 	fmt.Scan(&withdrawAmount)

		// 	if withdrawAmount <= 0 {
		// 		fmt.Println("Invalid amount. Must be graeter than 0")
		// 		continue
		// 	}

		// 	if withdrawAmount > accountBalance {
		// 		fmt.Println("You don't have enough balance to withdraw")
		// 		continue
		// 	}

		// 	accountBalance -= withdrawAmount
		// 	fmt.Println("Balance updated! New amount:", accountBalance)
		// } else {
		// 	fmt.Println("Goodbye!")
		// 	break
		// }
	}

}
