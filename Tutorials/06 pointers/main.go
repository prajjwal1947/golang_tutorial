package main

import "fmt"

func main() {
	fmt.Println("welcome to pointers")

	// var ptr *int
	// fmt.Println("value of ptr is", ptr)

	myNumber := 23
	var ptr = &myNumber // refrence means &

	fmt.Println("value of ptr is", ptr)
	fmt.Println("value of ptr is", *ptr)

	*ptr = *ptr + 2
	fmt.Println("new number is", myNumber)

}
