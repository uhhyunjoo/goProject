package main

import "fmt"

func main() {
	// var 변수명 [요소개수]타입
	day := [3]string{"mon", "tue", "wed"}

	for i := 0; i < 3; i++ {
		fmt.Println(day[i])
	}

	fmt.Println()

	var temps [2]float64 = [2]float64{24.3, 22}

	for i := 0; i < len(temps); i++ {
		fmt.Println(temps[i])
	}

	fmt.Println()

	var s = [5]int{1: 10, 3: 30}

	for i := 0; i < len(s); i++ {
		fmt.Println(s[i])
	}

	fmt.Println()

	x := [...]int{10, 20, 30, 40, 50}

	for i := 0; i < len(x); i++ {
		fmt.Println(x[i])
	}

	/// 흠냐
	source := []int{0, 1, 2}
	target := make([]int, len(source), cap(source)*2)
	copy(target, source)
	fmt.Println(target)               // [0 1 2 ] 출력
	println(len(target), cap(target)) // 3, 6 출력

	/// 흠냐

	// lst := make([]int, 5, cap(source)*2)
	// copy(target, source)
	// fmt.Println(target)               // [0 1 2 ] 출력
	// println(len(target), cap(target)) // 3, 6 출력

}
