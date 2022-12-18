package main

import (
	"fmt"
)

type User struct {
	Name string
	Age  int
}

func main() {
	myUser := User{"YC", 39}

	incrementAge(myUser)

	fmt.Println("========================")

	fmt.Println(myUser)
	fmt.Println(myUser.Name)
	fmt.Println(myUser.Age)

	fmt.Println("========================")

	incrementAgeReference(&myUser)
	fmt.Println(myUser)
	fmt.Println(myUser.Name)
	fmt.Println(myUser.Age)

	fmt.Println("========================")
	fmt.Println(myUser.prettyString())

}

func incrementAge(user User) {
	user.Age++
	fmt.Println(user.Age)
}

func incrementAgeReference(user *User) {
	user.Age++
	fmt.Println(user.Age)
}

func (user User) prettyString() string {
	return fmt.Sprintf("Pretty string %s", user.Name)
}
