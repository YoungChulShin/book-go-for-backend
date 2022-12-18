package main

import (
	"fmt"
)

func valueOfPi(multiplier uint) float32 {
	return 3.141592 * float32(multiplier)
}

func main() {
	pi := valueOfPi(2)
	fmt.Println(pi)

	fmt.Println("===========================")

	userName, age := nameAndAge(0)
	fmt.Println(userName, ", ", age)

	fmt.Println("===========================")

	var userName2 string
	var age2 int
	userName2, age2 = nameAndAge(2)
	fmt.Println(userName2, ", ", age2)

	fmt.Println("===========================")

	username3, _ := nameAndAge(3)
	fmt.Println(username3)

	fmt.Println("===========================")

	a, b := 1, 2
	fmt.Println(runMathOps(a, b, add))
	fmt.Println(runMathOps(a, b, sub))
	fmt.Println(runMathOps(a, b, mul))
	fmt.Println(runMathOps(a, b, div))

	fmt.Println("===========================")
	deferTest(5)
	fmt.Println("===========================")
	deferInlineTest(5)
}

func nameAndAge(userId uint) (string, int) {
	switch userId {
	case 0:
		return "YC", 39
	case 1:
		return "YE", 36
	case 2:
		return "MJ", 34
	default:
		return "NotFound", -1
	}
}

func runMathOps(a int, b int, op func(int, int) int) int {
	return op(a, b)
}

func add(a int, b int) int { return a + b }
func sub(a int, b int) int { return a - b }
func mul(a int, b int) int { return a * b }
func div(a int, b int) int { return a / b }

// defer가 붙으면 마지막 결과를 반환하기 전에 결과를 실행한다
func deferTest(x int) int {
	defer fmt.Println(x)
	y := x + 1
	fmt.Println(y)
	return y
}

func deferInlineTest(x int) int {
	defer func() {
		fmt.Println("defer start")
		z := x - 1
		fmt.Println(z)
		fmt.Println("defer end")
	}()
	y := x + 1
	fmt.Println(y)
	return y
}
