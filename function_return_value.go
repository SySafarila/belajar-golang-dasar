package main

import "fmt"

func getHello(name string) string {
	return "Hello " + name
}

func main() {
	result := getHello("Syahrul")
	fmt.Println(result)

	fmt.Println(getHello("Safarila"))
}
