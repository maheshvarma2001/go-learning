package main

import (
	"course/structs/user"
	"fmt"
)

func main() {
	userFirstName := getUserData("Please Enter your first name: ")
	userLastName := getUserData("Please Enter your last name: ")
	userBirthDate := getUserData("Please Enter your birthdate (MM/DD?YYYY): ")

	var appUser *user.User

	appUser, err := user.New(userFirstName, userLastName, userBirthDate)

	if err != nil {
		fmt.Println(err)
		return
	}

	appUser.OutputUserDetails()
	appUser.ClearUserName()
	appUser.OutputUserDetails()

}

func getUserData(promtText string) string {
	fmt.Print(promtText)
	var value string
	fmt.Scanln(&value) // Scanln to accept empty lines
	return value
}
