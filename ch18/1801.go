package main

import "fmt"

func main() {

	// var array [10]int

	var slice []int
	if len(slice) == 0 {
		fmt.Println("empty", slice)
	}

	slice[1] = 10 // slice 의 길이가 0인데 인덱스 1에 접근해서 에러 발생 (패닉)
	fmt.Println(slice[1])
}
