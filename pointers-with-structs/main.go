package main

import "fmt"

type User struct {
	name    string
	balance float64
}

func (user *User) move() {
	user.name = user.name + " Augusto"
	fmt.Printf("User %v moved\n", user.name)
}

func (user *User) print_name() {
	fmt.Printf("Current Name: %v\n", user.name)
}

func main() {
	user := &User{
		name: "Gustavo",
	}

	user2 := &User{
		name: "Marcus",
	}

	user.move()
	user2.move()

	// the name will not be updated beyond if you wont use pointers

	user.print_name()
	user2.print_name()
}
