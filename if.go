package main

import "fmt"

func main() {
	username := "SySafarila"

	if username == "SySafarila" {
		fmt.Println("Hello SySafarila")
	} else if username == "Syahrul" {
		fmt.Println("Hello Syahrul")
	} else {
		fmt.Println("Hello, may i know your username?")
	}

	if length := len(username); length > 5 {
		fmt.Println("Nama terlalu panjang")
	} else {
		fmt.Println("Nama sudah benar")
	}
}
