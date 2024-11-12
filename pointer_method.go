package main

import "fmt"

type Man struct {
	Name string
}

func (man *Man) Married() {
	man.Name = "Mr. " + man.Name
}

func main() {
	syahrul := Man{"Syahrul"}
	syahrul.Married()

	fmt.Println(syahrul)
}
