package main

import (
	"fmt"
)

func main() {
	//This is an exampe of redeclaration(we have one variable that is reassigned and one variable that is created.)
	var name string
	name, age := "Michael", 32
	fmt.Println(name, age)

	status := false

	michael, nico, status := 1, 2, true
	fmt.Println(michael, nico, status)
}
