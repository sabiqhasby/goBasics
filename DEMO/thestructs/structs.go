package main

import (
	"fmt"

	"example.com/structs/user"
)

func main() {
	userFirstName := user.GetUserData("Please enter your first name: ")
	userLastName := user.GetUserData("Please enter your last name: ")
	userBirthdate := user.GetUserData("Please enter your birthdate (MM/DD/YYYY): ")

	var appUser *user.User

	appUser, err := user.NewUser(userFirstName, userLastName, userBirthdate)
	if err != nil {
		fmt.Println(err)
		return
	}

	admin := user.NewAdmin("test@example.com", "test123")

	admin.OutputUserDetails()
	admin.ClearUserName()
	admin.OutputUserDetails()

	// ... do something awesome with that gathered data!
	appUser.OutputUserDetails()
	appUser.ClearUserName()
	appUser.OutputUserDetails()

}
