package main

import (
	"errors"
	"fmt"
	"time"
)

type user struct {
	firstName string
	lastName  string
	birthDate string
	createdAt time.Time
}

func (u *user) outputUserDetails() {

	fmt.Println(u.firstName, u.lastName, u.birthDate)
}

func (u *user) clearUserName() {
	u.firstName = ""
	u.lastName = ""
}

func newUser(firstName, lastName, birthDate string) (*user, error) {

	if firstName == "" || lastName == "" || birthDate == "" {
		return nil, errors.New("first name , last name and birth date are required")
	}

	return &user{
		firstName: firstName,
		lastName:  lastName,
		birthDate: birthDate,
		createdAt: time.Now(),
	}, nil
}

func main() {
	userFirstName := getUserData("Please Enter your first name: ")
	userLastName := getUserData("Please Enter your last name: ")
	userBirthDate := getUserData("Please Enter your birthdate (MM/DD?YYYY): ")

	var appUser *user

	appUser, err := newUser(userFirstName, userLastName, userBirthDate)

	if err != nil {
		fmt.Println(err)
		return
	}

	appUser.outputUserDetails()
	appUser.clearUserName()
	appUser.outputUserDetails()

}

func getUserData(promtText string) string {
	fmt.Print(promtText)
	var value string
	fmt.Scanln(&value) // Scanln to accept empty lines
	return value
}
