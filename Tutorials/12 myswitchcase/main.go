package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Hello, switch case in golang")

	rand.Seed(time.Now().UnixNano())
	diceNumber := rand.Intn(6) + 1
	fmt.Println(diceNumber)

	switch diceNumber {
	case 1:
		fmt.Print("dice value is 1 and you can open")
	case 2:
		fmt.Print("dice value is 2 and you can open")
	case 3:
		fmt.Print("dice value is 3 and you can open")
	case 4:
		fmt.Print("dice value is 4 and you can open")
	case 5:
		fmt.Print("dice value is 5 and you can open")
	case 6:
		fmt.Print("dice value is 6 and you can open")
	}

}
