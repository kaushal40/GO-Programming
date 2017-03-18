// struct
package main

import (
	"fmt"
)

func main() {
	type person struct {
		fname string
		lname string
		age   int
	}

	// you can declare new variable using below two ways
	// var newPerson person
	// newPerson :=  new(person)

	NewPerson := person{
		fname: "kaushal",
		lname: "prajapati",
		age:   24,
	}

	fmt.Println(NewPerson)
	fmt.Println(NewPerson.fname)

}
