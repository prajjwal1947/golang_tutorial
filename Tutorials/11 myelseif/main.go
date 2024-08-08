package main

import "fmt"

func main() {
	fmt.Print("welcome to if else")

	loginCount := 23
	var result string
	// if else
	// if else
	if loginCount < 10 {
		result = "regular user"
		fmt.Println("less than 10")
	} else if loginCount > 10 {
		result = "watch out"
		fmt.Println("greater than 10")
	} else {
		result = "exactly 10"
		fmt.Println("equal to 10")
	}
	fmt.Print(result)
}
