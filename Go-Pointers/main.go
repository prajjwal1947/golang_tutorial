package main

import "fmt"

// when we need to update the state
//pointers being 8 bytes
type User struct {
	email    string
	userName string
	age      int
}

func (u User) SetEmail(email string) {
	u.email = email
}

func (u User) Email() string {
	return u.email
}

func Email(u *User) string {
	return u.userName
}

// func main() {
// 	user := User{
// 		email: "prajjwal8055@gmail.com",
// 	}
// 	user.SetEmail("nagraj@gmail.com")
// 	fmt.Print(user.Email())

// }

///////////////////////////////////////////////////////////////////////////////////////////////////////

// second example
func modified(a string) {
	a = "modified"
}

func afterModification(a *string) {

	fmt.Println("inside mod string :", a)
	*a = "modified"
}

func main() {
	a := "hello"

	fmt.Println("before:", a)
	modified(a)
	fmt.Println("after:", a)

	//lets modify it actudally
	fmt.Println("after modification the result will ")
	afterModification(&a)
	fmt.Println("after modifincation the value of a will be:", a)

}
