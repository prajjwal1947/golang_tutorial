package main

import "fmt"

func main() {

	fmt.Print("struct is go lang")
	// no inhertince us  no super parents
	hitesh := User{"hitesh", "hitesh@go.dev", true, 21}
	fmt.Println(hitesh)
	fmt.Printf("hitesh details are here : %+v \n", hitesh)
	hitesh.GetStatus()
	hitesh.GetEmail()
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func (u User) GetStatus() {
	fmt.Print("Is user active :", u.Status)
}

func (u User) GetEmail() {
	u.Email = "test@go.dev"
	fmt.Print(" user email :", u.Email)
}
