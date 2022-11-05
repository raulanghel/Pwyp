// Exercise 2
// Declare a struct type and create a value of this type.
// Declare a function that can change the value of some field in this struct type.
// Display the value before and after the call to your function.
package main

import "fmt"

type user struct {
	name  string
	email string
	age   uint
}

func main() {
	user1 := user{
		name:  "Raul Anghel",
		email: "raul.anghel@gmail.com",
		age:   47,
	}

	fmt.Printf("Age of user1 is %v\n", user1.age)
	updateUserAge(&user1, 48)
	fmt.Printf("Age of user1 is %v\n", user1.age)
}

func updateUserAge(user *user, age uint) {
	user.age = age
}
