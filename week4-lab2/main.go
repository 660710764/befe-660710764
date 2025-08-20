package main

import (
	"fmt"
)

func main() {
	// var name string = "banlu"
	var age int = 20

	email := "chimsing_b@silpakorn.edu"
	gpa := 3.00

	firstname, lastname := "Banlu", "Chimsing"

	fmt.Printf("Name %s %s, age %d, email %s, gpa%.2f\n", firstname, lastname, age, email, gpa)
}
