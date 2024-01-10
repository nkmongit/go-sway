package user

import (
	"errors"
	"fmt"
	"time"
)

// defining the struct
type User struct {
	firstName string
	lastName  string
	birthDate string
	createdAt time.Time
}

// struct embedding
type Admin struct {
	email    string
	password string
	User
}

// struct method, it belongs to the User struct
func (u User) OutputUserDetails() {
	fmt.Println(u.firstName, u.lastName, u.birthDate)
}

// reciever argument -> (u User)
// we are passing this reciever argument as pointer referencing to the struct
// because these just work like a variable and we do not want to create copies
func (u *User) ClearUserName() {
	u.firstName = ""
	u.lastName = ""
}

// constructor function for Admin
func NewAdmin(email, password string) Admin {
	return Admin{
		email:    email,
		password: password,
		User: User{
			firstName: "ADMIN",
			lastName:  "ADMIN",
			birthDate: "___",
			createdAt: time.Now(),
		},
	}
}

// constructor function - utility function that creates a struct
func New(firstName, lastName, birthDate string) (*User, error) {
	// validations
	if firstName == "" || lastName == "" || birthDate == "" {
		return nil, errors.New("First name, last name and birthdate are required")
	}
	return &User{
		firstName: firstName,
		lastName:  lastName,
		birthDate: birthDate,
		createdAt: time.Now(),
	}, nil
}
