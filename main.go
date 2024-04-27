package main

import "fmt"

func main() {
	fmt.Println("Hello world")

	// declare variable dengan manifest typing
	var firstName string = "oscar"
	var lastName string
	lastName = "isaac"

	fmt.Printf("hello %s %s!\n", firstName, lastName)
}
