package main

import "fmt"

func main() {
	var apple int
	var orange int32
	//int and int32 are not equal
	apple = int(orange)
	_ = apple
	//This will print ascii characters.
	orange = 65
	color := string(orange)
	fmt.Println(color)
	//another way to do ascii characters
	fmt.Println(string([]byte{104, 105}))

}
