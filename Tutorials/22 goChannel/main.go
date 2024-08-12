package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()

	defer func() {
		fmt.Println(time.Since(now))
	}()

	smokeSignal := make(chan bool)
	eveilNinja := "Tommy,"
	go attack(eveilNinja, smokeSignal)
	// smokeSignal <- false deadlock will occurs
	fmt.Println(<-smokeSignal)

}

func attack(target string, attacked chan bool) {
	time.Sleep(time.Second)

	fmt.Println("throwing ninja star", target)

	attacked <- true
}
