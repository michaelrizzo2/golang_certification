package main

import (
	"fmt"
)

func main() {
	safe := true
	fmt.Println(safe)

	//var name type=value is one way
	//another way is to have go guess the type from the output var name=value
	//the final way is name:=value this is short declaration.

	var name = "carl"
	var isscientist = true
	var age = 62
	var degree = 5.
	fmt.Println(name, isscientist, age, degree)
}
