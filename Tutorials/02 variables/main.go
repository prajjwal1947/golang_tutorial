package main

import "fmt"

const Login string = "hdhkkhoho" // public

func main() {
	var username string = "hitesh"
	fmt.Println(username)
	fmt.Printf("variables type : %T \n", username)

	var isloggedIn bool = false
	fmt.Println(isloggedIn)
	fmt.Printf("variables type : %T \n", isloggedIn)

	var smallVal uint8 = 255
	fmt.Println(smallVal)
	fmt.Printf("variables type : %T \n", smallVal)

	var fkoatVal float32 = 255.3686683793
	fmt.Println(fkoatVal)
	fmt.Printf("variables type : %T \n", fkoatVal)

	// deafult value and aliases

	var anotherVariable int
	fmt.Printf("variables type : %T \n", anotherVariable)

	//implicit type
	var website = "prajjwal"
	fmt.Println(website)

	// no var style
	numberOfUser := 100000.0
	fmt.Println(numberOfUser)

	//login token

	fmt.Println(Login)

}
