package main

import (
	"fmt"
)

func main() {
	// index를 가져오기
	names := []string{"YC", "YE", "MJ"}
	for i := range names {
		fmt.Println(names[i])
	}

	fmt.Println("========================")

	// index, value 를 가져오기
	for i, v := range names {
		fmt.Println(i)
		fmt.Println(v)
	}

	fmt.Println("========================")

	// index상관 없이 value 를 가져오기
	for _, v := range names {
		fmt.Println(v)
	}

	fmt.Println("========================")

	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	fmt.Println("========================")

	for i := 0; i < len(names); i++ {
		fmt.Println(i)
	}
}
