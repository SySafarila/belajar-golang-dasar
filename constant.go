package main

import "fmt"

func main() {
	const (
		firstName string = "Syahrul"
		lastName         = "Safarila"
	)

	fmt.Println(firstName)
	fmt.Println(lastName)

	// error
	// firstName = "SySafarila"
	// lastName = "SafarilaSy"
}
