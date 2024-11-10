package main

import "fmt"

func main() {
	counter := 1

	for counter <= 10 {
		fmt.Println("Perulangan ke ", counter)
		counter++
	}

	fmt.Println("Selesai")

	for i := 1; i <= 10; i++ {
		fmt.Println("Perulangan ke ", i)
	}

	fmt.Println("Selesai")

	names := []string{"Syahrul", "Safarila", "SySafarila"}

	for index, name := range names {
		fmt.Println("index", index, "=", name)
	}

	for _, name := range names {
		fmt.Println(name)
	}
}
