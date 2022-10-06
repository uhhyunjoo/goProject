package main

import "fmt"

func main() {
	// var 변수명 [요소개수]타입
	a := [...]int{1, 2, 3, 4, 5}
	b := [...]int{55, 44, 33, 22, 11}

	for i, v := range a {
		fmt.Printf("a[%d] = %d\n", i, v)
	}
	fmt.Println()

	for i, v := range b {
		fmt.Printf("b[%d] = %d\n", i, v)
	}

	b = a // 그냥 복사 (서로 길이 다른 경우 안됨 ㅠ)

	fmt.Println()
	for i, v := range b {
		fmt.Printf("b[%d] = %d\n", i, v)
	}

}
