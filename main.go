package main

func main() {
	// Loop over collection
	slice := []int{1, 2, 3}

	for i, v := range slice {
		println(i, v)
	}
}
