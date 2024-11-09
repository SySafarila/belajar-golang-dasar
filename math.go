package main

import "fmt"

func main() {
	var a = 10
	var b = 10
	var c = a + b
	fmt.Println(c)

	var i = 10
	i += 10 // i = i + 10
	fmt.Println(i)
	i /= 2 // i = i / 2
	fmt.Println(i)

	var j = 1
	j++ // j = j + 1
	fmt.Println(j)
	j++ // j = j + 1
	fmt.Println(j)
	j-- // j = j - 1
	fmt.Println(j)
}
