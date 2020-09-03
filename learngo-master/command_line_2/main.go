package main

import ("fmt";"os")

func main(){
	fmt.Printf("%#v\n",os.Args)
	fmt.Println("Path is ",os.Args[0])
	fmt.Println("1st arguement is ",os.Args[1])
	fmt.Println("2nd arguement is ",os.Args[2])
	fmt.Println("3rd arguement is  ",os.Args[3])
	fmt.Println(len(os.Args))
}