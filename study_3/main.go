package main

import (
	"fmt"
)

type person struct {
	firstName string
	lastName  string
}

func main() {
	// nozomu := person{firstName:"Nozomu",lastName:"Tsuruta"}
	// fmt.Println(nozomu)

	var nozomu person

	nozomu.firstName = "Nozomu"
	nozomu.lastName = "Tsuruta"

	fmt.Println(nozomu)
	fmt.Printf("%+v", nozomu)
}
