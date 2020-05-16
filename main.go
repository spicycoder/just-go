package main

import (
	"fmt"
)

func main() {
	sum, diff := calculate(10, 5)
	fmt.Println(sum)
	fmt.Println(diff)
}

func calculate(x int, y int) (int, int) {
	return x + y, x - y
}
