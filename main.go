package main

import "fmt"

func main() {
	slice := []int{1, 2, 3, 4, 5, 6}
	fmt.Println(slice)
	fmt.Println(slice[:])
	fmt.Println(slice[1:])
	fmt.Println(slice[:4])
	fmt.Println(slice[2:3])
}
