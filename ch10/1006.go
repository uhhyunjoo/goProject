package main

import "fmt"

var age int = 100

func getMyAge() int {
	return 25
}

func main() {

	fmt.Println("age is", age)

	age := 10

	switch age := getMyAge(); age {
	case 10:
		fmt.Println("Teen")
	case 33:
		fmt.Println("Pair 3")
	default:
		fmt.Println("My age is", age)
	}

	// fmt.Println("age is", age) // switch 문 바깥이라 출력 안됨
	fmt.Println("age is", age) // switch 문 바깥이라 출력 안됨 (근데 얘는 위에서 했으니깐 됨 ㅇㅇ)
}
