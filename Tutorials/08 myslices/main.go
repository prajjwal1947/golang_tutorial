package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Print("welcome to slices")

	var fruitList = []string{"apple", "tomato", "peches"}

	fmt.Printf("type of fruitlist %T\n", fruitList)

	fruitList = append(fruitList, "mango")
	fmt.Println(fruitList)

	fruitList = append(fruitList[1:3])
	fmt.Println(fruitList)

	highScores := make([]int, 4)
	highScores[0] = 234
	highScores[1] = 34
	highScores[2] = 534
	highScores[3] = 834
	// highScores[4]=934;

	highScores = append(highScores, 555, 999)

	fmt.Println(highScores)

	sort.Ints(highScores)
	fmt.Println(highScores)

	// how to remove a value from a slice base on index

	var courses = []string{"reactjs", "angular", "vue", "nodejs"}
	fmt.Println(courses)
	var index int = 2
	courses = append(courses[:index], courses[index+1:]...)
	fmt.Println(courses)

}
