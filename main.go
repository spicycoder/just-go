package main

import "fmt"

func main() {
	firstName := "John"
	firstNamePointer := &firstName

	fmt.Println(firstNamePointer)
	fmt.Println(*firstNamePointer)
	fmt.Println(&firstNamePointer)
}
