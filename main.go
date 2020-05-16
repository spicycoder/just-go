package main

import (
	"fmt"

	"github.com/spicycoder/just-go/models"
)

func main() {
	u := models.User{
		ID:        1,
		FirstName: "John",
		LastName:  "Doe",
	}

	fmt.Println(u)
}
