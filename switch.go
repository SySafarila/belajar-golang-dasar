package main

import "fmt"

func main() {
	name := "Safarilax"

	switch name {
	case "Syahrul":
		fmt.Println("Hello Syahrul")
	case "Safarila":
		fmt.Println("Hello Safarila")
	default:
		fmt.Println("Hello Tak Dikenal")
	}

	switch length := len(name); length > 5 {
	case true:
		fmt.Println("Nama terlalu panjang")
	case false:
		fmt.Println("Nama sudah benar")
	}

	length := len(name)

	switch {
	case length > 5:
		fmt.Println("Nama terlalu panjang")
	case length > 5:
		fmt.Println("Nama pas")
	default:
		fmt.Println("Nama default")
	}
}
