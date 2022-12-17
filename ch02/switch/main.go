package main

import (
	"fmt"
)

func main() {
	name := "NewBie"

	switch name {
	case "YE":
		fmt.Println("Hi, YE")
	case "MJ":
		fmt.Println("Hi, MJ")
	case "YC":
		fmt.Println("Hi, YC")
	default:
		fmt.Println("Can't find. ", name)
	}
}
