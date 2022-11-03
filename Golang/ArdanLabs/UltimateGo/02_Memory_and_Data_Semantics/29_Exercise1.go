// Part A: Declare a struct type to maintain information about a user (name, email and age).
// Create a value of this type, initialize with values and display each field.

// Part B: Declare and initialize an anonymous struct type with the same three fields. Display the value.
package main

import "fmt"

type user struct {
	name  string
	email string
	age   uint
}

func main() {
	user1 := user{
		name:  "Raul",
		email: "raul.anghel@gmaul.com",
		age:   48,
	}

	fmt.Printf("%+v\n", user1)

	user2 := user{}
	user2.name = "Vlad"
	user2.email = "vlad.anghel@gmail.com"
	user2.age = 5

	fmt.Printf("%+v\n", user2)

	user3 := struct {
		name  string
		email string
		age   uint
	}{
		name:  "Sabina",
		email: "sabina.anghel@gmail.com",
		age:   5,
	}

	fmt.Printf("%+v\n", user3)
}
