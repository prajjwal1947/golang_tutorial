package main

import "fmt"

func main() {
	fmt.Println("welcome to go-lang array")

	var fruitList [4]string

	fruitList[0] = "Apple"
	fruitList[1] = "Mango"
	fruitList[2] = "Banana"

	fmt.Println("fruitList list is:", fruitList)
	fmt.Println("fruitList list is:", len(fruitList))

	var vegList = [3]string{"Potato", "Tomato", "Onion"}
	fmt.Println("vegList list is:", vegList)
	fmt.Println("vegList list is:", len(vegList))
}
