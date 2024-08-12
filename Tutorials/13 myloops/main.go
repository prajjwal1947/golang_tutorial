package main

import "fmt"

func main() {
	fmt.Printf("looops i go lang ")

	days := []string{"sunday", "monday", "tuesday", "wednesday", "thursday", "friday", "saturday"}

	fmt.Println(days)

	//fist way

	// for d := 0; d < len(days); d++ {
	// 	fmt.Println(days[d])
	// }

	//second way

	for i := range days {
		fmt.Println(days[i])
	}

	// thid way

	for index, day := range days {
		fmt.Printf("%v %v \n", index, day)
	}

	//while loop
	rogueValue := 1

	for rogueValue < 10 {
		if rogueValue == 5 {
			// break
			rogueValue++
			continue
		}
		fmt.Println("value is ", rogueValue)
		rogueValue++
	}
}
