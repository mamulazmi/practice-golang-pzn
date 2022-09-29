package main

import "fmt"

func main() {
	const firstName string = "Azmi"
	const lastName = "Imamul"
	const pi = 3.14

	fmt.Println("FirstName :", firstName)
	fmt.Println("LastName :", lastName)
	fmt.Println("Pi :", pi)

	const (
		age     int = 10
		country     = "Indonesia"
	)

	fmt.Println(age)
	fmt.Println(country)
}
