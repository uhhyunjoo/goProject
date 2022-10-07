package main

import "fmt"

func main() {

	// (1) {} 으로 초기화 하기 ([]안에 길이 넣지 x)
	// var slice []int{1,2,3}

	// (2) make() 로 초기화
	// var slice = make([]int, 3) // 길이가 3인 int 슬라이스값, 각 슬라이스 요솟값은 int의 기본 타입인 0

	// var array = [...]int{1,2,3} // 배열
	// var slice = []int{1,2,3} // 슬라이스

	//var slice1 []int{1, 2, 3} //뭐지 왜 안되지

	var slice []int
	slice = []int{1, 2, 3}

	// var slice3 []int = []int{1, 2, 3, 4, 5}

	// var slice4 = make([]int, 3)

	//var slice = make([]int, 3)

	slice2 := append(slice, 4)

	fmt.Println(slice)
	fmt.Println(slice2)

	slice3 := append(slice2, 5, 6, 7)

	fmt.Println(slice3)
}
