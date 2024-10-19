package main

import (
	"fmt"
	"time"
)

type user struct {
	firstName string
	lastName  string
	birthDate string
	createdAt time.Time
}

func (u user) outputUserDetails() {

	fmt.Println(u.firstName, u.lastName, u.birthDate)
}

func main() {
	userFirstName := getUserData("Please Enter your first name: ")
	userLastName := getUserData("Please Enter your last name: ")
	userBirthDate := getUserData("Please Enter your birthdate (MM/DD?YYYY): ")

	var appUser user

	appUser = user{
		firstName: userFirstName,
		lastName:  userLastName,
		birthDate: userBirthDate,
		createdAt: time.Now(),
	}

	outputUserDetails(&appUser)
}

func outputUserDetails(u *user) {
	fmt.Println(u.firstName, u.lastName, u.birthDate)
}

func getUserData(promtText string) string {
	fmt.Print(promtText)
	var value string
	fmt.Scan(&value)
	return value
}
