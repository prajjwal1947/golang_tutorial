package main

import "fmt"

func main() {
	fmt.Print("welcome to function in golang")
	greater()
	result := adder(3, 5)
	fmt.Println(result)
	proResult := proAdder(2, 4, 6, 8, 10)
	fmt.Print(proResult)
}

func greater() {
	fmt.Print("namaster go lang")
}

func proAdder(values ...int) int {
	total := 0
	for _, val := range values {
		total += val
	}
	return total
}

func adder(val1 int, val2 int) int {
	return val1 + val2
}
