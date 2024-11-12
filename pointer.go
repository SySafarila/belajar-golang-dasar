package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func main() {
	address1 := Address{"Subang", "Jawa Barat", "Indonesia"}
	address2 := &address1 // pointer atau alias

	address2.City = "Bandung"
	fmt.Println(address1)
	fmt.Println(address2)
}
