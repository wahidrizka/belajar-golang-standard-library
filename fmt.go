package main

import "fmt"

func main() {
	firstName := "Wahid"
	lastName := "Fathurrohman"

	fmt.Println("Hello '", firstName, lastName, "'")

	fmt.Printf("Hello '%s %s'\n", firstName, lastName)
}
