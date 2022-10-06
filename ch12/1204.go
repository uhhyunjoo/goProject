package main

import "fmt"

func main() {
	// var 변수명 [요소개수]타입
	t := [...]int{10, 20, 30, 40, 50}
	t[2] = 300

	for i, v := range t {
		fmt.Println(i, v)
	}

	a, b := 10, 20

	fmt.Println(a)
	fmt.Println(b)

	// c, d := range t
	// 왜 안 받아지지 ㅠ

	// fmt.Println(c)
	// fmt.Println(d)
}
