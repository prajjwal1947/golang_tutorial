package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Hello World")

	start := time.Now()

	defer func() {
		fmt.Print(time.Since(start))
	}()

	evilNinjas := []string{"John", "Paul", "George", "Ringo"}
	for _, evilNinjas := range evilNinjas {
		go attack(evilNinjas)
	}
	time.Sleep(time.Second * 2)

}

func attack(target string) {
	fmt.Println("Throwing ninja targets at ", target)
	time.Sleep(time.Second)
}
