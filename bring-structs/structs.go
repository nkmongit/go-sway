package main

import (
	"fmt"

	"example.com/structs/user"
)

func main() {
	userFirstName := getUserData("Please enter your first name: ")
	userLastName := getUserData("Please enter your last name: ")
	userBirthDate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")

	// creating the struct variable of type User
	var appUser *user.User
	// instantiating appUser with User struct
	appUser, err := user.New(userFirstName, userLastName, userBirthDate)
	if err != nil {
		fmt.Println(err)
		return
	}
	// we can also do like this as we
	// are writing all the data in order
	// appUser = User{
	//  userFirstName,
	//	userLastName,
	//	userBirthDate,
	//	time.Now(),
	// }

	// making use of Admin struct
	admin := user.NewAdmin("test@example.com", "example")
	admin.OutputUserDetails()
	admin.ClearUserName()
	admin.ClearUserName()

	appUser.OutputUserDetails()
	appUser.ClearUserName()
	appUser.OutputUserDetails()
}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	// fmt.Scan(&value)
	// Scanln, lets you press ENTER key and take that as value
	fmt.Scanln(&value)
	return value
}
