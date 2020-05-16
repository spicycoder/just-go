package main

import "fmt"

func main() {
	type user struct {
		ID        int
		FirstName string
		LastName  string
	}

	var u user
	u = user{
		ID:        1,
		FirstName: "John",
		LastName:  "Doe",
	}

	fmt.Println(u)
	fmt.Println(u.FirstName)
}
