package main

import "fmt"

func main() {

	channel := make(chan string, 2)

	channel <- "first message "
	channel <- "second message " // deadlock

	fmt.Println(<-channel)
	fmt.Println(<-channel)
	// fmt.Println(<-channel)
	fmt.Println(<-channel)

}
