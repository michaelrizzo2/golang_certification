package main

import (
	"fmt"
)

func main() {
	//This wil be modified with the new multiple variable syntax.
	var (
		speed int
		heat  float64
		off   bool
		brand string
	)
	fmt.Println(speed, heat, off, brand)
	//string variables default to ""
	// int and floats default ot zero and bolean values default to false.
}
