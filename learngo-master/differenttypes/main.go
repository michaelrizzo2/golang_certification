package main

import "fmt"

func main() {
	var apple int
	var orange int32
	//int and int32 are not equal
	apple = int(orange)
	fmt.Println(apple, orange)
}
