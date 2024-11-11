package main

import "fmt"

type Customer struct {
	Name    string
	Address string
	Age     int
}

func (customer Customer) sayHello(name string) {
	fmt.Println("Hello, my name is", customer.Name, "Nice to meet you", name)
}

func main() {
	var syahrul Customer
	syahrul.Name = "Syahrul Safarila"
	syahrul.Address = "Cianjur"
	syahrul.Age = 23

	fmt.Println(syahrul)
	fmt.Println(syahrul.Name)
	fmt.Println(syahrul.Address)
	fmt.Println(syahrul.Age)

	susi := Customer{
		Name:    "Susilawati",
		Address: "Lembang",
		Age:     22,
	}
	fmt.Println(susi)
	susi.sayHello("Syahrul")

	budi := Customer{"Budi", "Indonesia", 20}
	fmt.Println(budi)
}
