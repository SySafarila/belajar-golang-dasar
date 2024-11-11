package main

import "fmt"

func getFullName() (string, string) {
	return "Syahrul", "safarila"
}

func main() {
	// firstName, lastName := getFullName()
	// fmt.Println(firstName, lastName)

	// gunakan _ kalo gamau pake lastName nya
	firstName, _ := getFullName()
	fmt.Println(firstName)
}
