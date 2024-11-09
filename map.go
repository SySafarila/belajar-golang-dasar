package main

import "fmt"

func main() {
	person := map[string]string{
		"name":    "Syahrul",
		"address": "Cianjur",
	}

	fmt.Println(person)
	fmt.Println(person["name"])
	fmt.Println(person["address"])
}
