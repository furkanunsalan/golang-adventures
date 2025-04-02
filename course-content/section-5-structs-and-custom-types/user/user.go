package user

import (
	"errors"
	"fmt"
	"time"
)

// Fields of outside accesible types should also be uppercase to export properly
type User struct {
	firstName string
	lastName  string
	birthdate string
	createdAt time.Time
}

// struct embedding: admin user has every field of User + new ones
type Admin struct {
	email    string
	password string
	User // explicit struct embedding
	// User User -> implicit struct embedding, do not prefer if you dont want to do admin.User.outputUserDetails()
}

// Paranthesis mean "receiver", which connect the function to a struct
func (u User) OutputUserDetails() {
	// (*u).firstName -> better way of dereferencing and using variables, but mostly not used because go allows the following way
	fmt.Println(u.firstName, u.lastName, u.birthdate)
	fmt.Println("Created at: ", u.createdAt)
}

// You must accept a pointer if your dunction is a mutation function, you can also make that for other methods but it's not necessary
func (u *User) ClearUsername() {
	u.firstName = ""
	u.lastName = ""
}

func NewAdmin(email, password string) Admin {
	return Admin{
		email:    email,
		password: password,
		User: User{
			firstName: "ADMIN",
			lastName:  "ADMIN",
			birthdate: "---",
			createdAt: time.Now(),
		},
	}
}

// You can add validation when you create instances using constructor functions
func New(firstname, lastname, birthdate string) (*User, error) {
	if firstname == "" || lastname == "" || birthdate == "" {
		return nil, errors.New("user information can't be empty")
	}

	return &User{
		firstName: firstname,
		lastName:  lastname,
		birthdate: birthdate,
		createdAt: time.Now(),
	}, nil

}
