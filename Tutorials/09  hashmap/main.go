package main

import "fmt"

func main() {
	fmt.Printf("welcome to hashmap")

	lauguages := make(map[string]string)

	lauguages["js"] = "java-script"
	lauguages["rb"] = "ruby"
	lauguages["py"] = "python"
	lauguages["go"] = "golang"

	fmt.Println(lauguages)

	// delete(lauguages, "rb")
	fmt.Println(lauguages)

	//loops for map

	for key, value := range lauguages {
		fmt.Printf("for key %v, value is %v\n", key, value)
	}

	//another way to declare map

}
