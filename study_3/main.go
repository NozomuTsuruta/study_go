package main

import (
	"fmt"
)

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	// contactInfo contactInfo
	contactInfo
}

func main() {
	nozomu := person{firstName: "Nozomu", lastName: "Tsuruta", contactInfo: contactInfo{email: "nozomu@gmail.com", zipCode: 12345}}

	// var nozomu person

	// nozomu.firstName = "Nozomu"
	// nozomu.lastName = "Tsuruta"

	// fmt.Println(nozomu)

	nozomuPointer := &nozomu
	nozomuPointer.updateName("Jonas")
	nozomu.print()
}

func (p *person) updateName(newFirstName string) {
	(*p).firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
