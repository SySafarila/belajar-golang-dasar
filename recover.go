package main

import "fmt"

func endApp() {
	fmt.Println("End app")
	message := recover()
	fmt.Println("Error", message)
}

func runApp(error bool) {
	defer endApp()

	if error {
		panic("Oops error gaes")
	}
}

func main() {
	runApp(!false)
}
