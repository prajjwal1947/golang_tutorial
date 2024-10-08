package main

import "fmt"

func main() {
	ninja1, ninja2 := make(chan string), make(chan string)

	go captainElect(ninja1, "ninja 1")
	go captainElect(ninja2, "ninja 2")

	// fmt.Println(<-ninja1)
	// fmt.Println(<-ninja2)
	select {
	case message := <-ninja1:
		fmt.Println(message)
	case message := <-ninja2:
		fmt.Println(message)
	}

	roughlyFair()

}

func captainElect(ninja chan string, message string) {
	ninja <- message
}

func roughlyFair() {
	ninja1 := make(chan interface{})
	close(ninja1)
	ninja2 := make(chan interface{})
	close(ninja2)

	var ninja1Count, ninja2Count int

	for i := 0; i < 1000; i++ {
		select {
		case <-ninja1:
			ninja1Count++
		case <-ninja2:
			ninja2Count++
		}
	}

	fmt.Printf("ninja-count: %d ,ninja2count :%d/n", ninja1Count, ninja2Count)
}
