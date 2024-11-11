package main

import "fmt"

type Filter func(string) string

func sayHelloWithFilter(name string, filter Filter) {
	filteredName := filter(name)

	fmt.Println("Hello ", filteredName)
}

func spamFIlter(name string) string {
	if name == "Anjing" {
		return "..."
	} else {
		return name
	}
}

func main() {
	sayHelloWithFilter("Syahrul", spamFIlter)

	filter := spamFIlter
	sayHelloWithFilter("Anjing", filter)
}
