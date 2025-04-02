package main

import (
	"fmt"
	"example.com/structs/user"
)

func main() {
	userFirstName := getUserData("Please enter your first name: ")
	userLastName := getUserData("Please enter your last name: ")
	userBirthdate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")

	// var appUser user = user{
	// 	firstName: userFirstName,
	// 	lastName:  userLastName,
	// 	birthdate: userBirthdate,
	// 	createdAt: time.Now(),

	// 	// Also can be done by ommitting the keys and passing the values in order
	// 	// userFirstName,
	// 	// userLastName,
	// 	// userBirthdate,
	// 	// time.Now(),
	// }

	// a better "convention" of creating instances of types with constructor functions
	var appUser *user.User 
	appUser, err := user.New(userFirstName, userLastName, userBirthdate)

	if err != nil {
		fmt.Println(err)
		return
	}

	admin := user.NewAdmin("test@example.com", "test123")
	
	// admin.User.OutputUserDetails() // if you define the user implicitly like User User
	admin.OutputUserDetails() // if you define the user explicitly like just User

	// ... do something awesome with that gathered data!

	// Call the method of the struct instance
	appUser.OutputUserDetails()

	appUser.ClearUsername() // a copy is created thats why it doesn't clear the username
	appUser.OutputUserDetails()
}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	// fmt.Scan(&value) // doesn't stop when you enter an empty line so you need to use the following for those cases
	fmt.Scanln(&value)
	return value
}
