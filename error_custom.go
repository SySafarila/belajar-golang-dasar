package main

import "fmt"

type validationError struct {
	Message string
}

func (v *validationError) Error() string {
	return v.Message
}

type notFoundError struct {
	Message string
}

func (n *notFoundError) Error() string {
	return n.Message
}

func SaveData(id string, data any) error {
	if id == "" {
		return &validationError{Message: "Validation error"}
	}

	if id != "syahrul" {
		return &notFoundError{Message: "Data not found"}
	}

	return nil
}

func main() {
	err := SaveData("syahrul", nil)

	if err != nil {
		// terjadi error
		if validationErr, ok := err.(*validationError); ok {
			fmt.Println("Validation error:", validationErr.Error())
		} else if notFoundErr, ok := err.(*notFoundError); ok {
			fmt.Println("Not found error:", notFoundErr.Error())
		} else {
			fmt.Println("Unknown error:", err.Error())
		}
	} else {
		// sukses
		fmt.Println("Sukses")
	}
}
