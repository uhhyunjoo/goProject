package main

import "fmt"

func main() {
	// var 변수명 [요소개수]타입
	a := [2][5]int{{1, 2, 3, 4, 5}, {11, 12, 13, 14, 15}}

	for _, arr := range a {
		for _, v := range arr {
			fmt.Print(v, " ")
		}
		fmt.Println()
	}

	// for i := a{
	// 	fmt.Println(i)
	// }
	// go lang 배열 순회할 떄 index 안 받고, index 로 접근 안하고 그냥 도는 방법은 없나?
}
