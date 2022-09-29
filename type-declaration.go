package main

import "fmt"

func main() {

	type NoKTP string

	var noKTP NoKTP = "123456789"
	fmt.Println("No KTP :", noKTP)

	type Married bool

	var marriedStatus Married = true
	fmt.Println(marriedStatus)

}
