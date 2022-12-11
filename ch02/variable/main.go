package main

import (
	"fmt"
)

// 전역 변수 선언
var age = 16

// 멀티 변수 선언
var (
	schoolName = "항공대"
	schoolAge  = 40
)

// 상수 선언
const name1 = "테스트 이름1"
const name2 string = "테스트 이름2"

func main() {
	fmt.Println(age)
	age = 20
	fmt.Println(age)

	// 내부 변수 선언
	var name = "테스트"
	fmt.Println(name)

	fmt.Println(schoolName)
	fmt.Println(schoolAge)

	fmt.Println(name1)
	fmt.Println(name2)

	// 지역 변수 선언 및 할당
	myName := "Anne"
	fmt.Println(myName)

	// 배열 선언
	names := []string{"name1", "name2", "name3"}
	names = append(names, "name4")
	fmt.Println(names)

	var names2 []string = []string{"name1", "name2", "name3"}
	names2[1] = "name2-1"
	fmt.Println(names2)

	// 고정 배열 선언
	fixNames := make([]string, 3)
	fixNames[0] = "hello"
	fixNames[1] = "world"
	fixNames[2] = "go"
	fmt.Println(fixNames)

	fixNames = append(fixNames, "newbie")
	fmt.Println(fixNames)

	// 고정 배열 + 초기 값
	// 처음에는 0사이즈의 배열을 만들지만, 추가로 3까지는 용량이 확보되어 있다
	// make([]string, 0, 3)
}
