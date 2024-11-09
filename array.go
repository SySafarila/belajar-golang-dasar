package main

import "fmt"

func main() {
	// fixed array
	var names [3]string
	names[0] = "Syahrul"
	names[1] = "(SySafarila)"
	names[2] = "Safarila"

	fmt.Println(names[0])
	fmt.Println(names[1])
	fmt.Println(names[2])

	var values = [3]int{90, 80, 50}
	fmt.Println(values)

	// dynamic array
	var members = [...]string{
		"syahrul",
		"safarila",
		"sysafarila",
		"susi",
	}
	fmt.Println(members)
	fmt.Println(len(members))
}
